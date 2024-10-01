package data

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ciphermountain/deadenz/pkg/opts"
)

type Parser func([]byte, any) error

type Loader interface {
	Data(context.Context) ([]byte, error)
}

type DataLoader struct {
	mu    sync.RWMutex
	cache *dataCache

	configs   map[string]*loadConfig
	reloaders map[string]reloadConfig

	starter sync.Once
	closer  sync.Once
	running atomic.Bool
	chClose chan struct{}
}

func NewDataLoader() *DataLoader {
	return &DataLoader{
		cache:     newDataCache(),
		configs:   make(map[string]*loadConfig),
		reloaders: make(map[string]reloadConfig, 0),
		chClose:   make(chan struct{}, 1),
	}
}

func (l *DataLoader) SetLoader(typeKey reflect.Type, loader Loader, parser Parser, options ...opts.Option) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	language := &language{lang: ""}
	config := &loadConfig{
		valueType: typeKey,
		parser:    parser,
		loader:    loader,
		cancel:    func() {},
	}

	for _, opt := range options {
		opt(config)
		opt(language)
	}

	key := makeCacheKey(typeKey, language.lang)

	for _, opt := range options {
		opt(setReloadConfig(key, l.reloaders))
	}

	l.configs[key] = config

	return nil
}

func (l *DataLoader) Load(value any, options ...opts.Option) error {
	return l.LoadCtx(context.Background(), value, options...)
}

func (l *DataLoader) LoadCtx(ctx context.Context, returnVal any, options ...opts.Option) error {
	language := &language{lang: ""}
	for _, opt := range options {
		opt(language)
	}

	value, err := valueFrom(returnVal)
	if err != nil {
		return err
	}

	key := makeCacheKey(value.Type(), language.lang)

	if cachedValue, exists := l.cache.get(key); exists {
		// value and cachedValue are already accessed via reflect.Indirect
		value.Set(cachedValue)

		return nil
	}

	return l.doLoad(ctx, key, returnVal)
}

func (l *DataLoader) doLoad(ctx context.Context, key string, returnVal any) error {
	l.mu.RLock()
	defer l.mu.RUnlock()

	config, exists := l.configs[key]
	if !exists || config.loader == nil {
		return fmt.Errorf("loader does not exist for %+v", reflect.TypeOf(returnVal))
	}

	if config.parser == nil {
		return fmt.Errorf("parser does not exist for %+v", reflect.TypeOf(returnVal))
	}

	bts, err := config.loader.Data(ctx)
	if err != nil {
		return err
	}

	if err := config.parser(bts, returnVal); err != nil {
		return err
	}

	// only store an indirect value type and not a pointer
	l.cache.set(key, reflect.Indirect(reflect.ValueOf(returnVal)))

	return nil
}

func (l *DataLoader) backgroundLoad(ctx context.Context, key string) error {
	l.mu.RLock()
	defer l.mu.RUnlock()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// call previous cancel to ensure cache calls remain clean
	// save new cancel call
	l.configs[key].cancel()
	l.configs[key].cancel = cancel

	conf, ok := l.configs[key]
	if !ok {
		return errors.New("no configuration for key")
	}

	retVal := reflect.New(conf.valueType)

	return l.doLoad(ctx, key, retVal.Interface())
}

func (l *DataLoader) Start() error {
	l.starter.Do(func() {
		go l.run()
		l.running.Store(true)
	})

	return nil
}

func (l *DataLoader) Close() error {
	if !l.running.Load() {
		return fmt.Errorf("service not running")
	}

	l.closer.Do(func() {
		close(l.chClose)
	})

	return nil
}

func (l *DataLoader) run() {
	ticker := time.NewTicker(250 * time.Millisecond)
	loadCtx, cancelAll := context.WithCancel(context.Background())

	for {
		select {
		case <-ticker.C:
			// search configs and run loaders
			l.mu.RLock()
			reloaders := l.reloaders
			l.mu.RUnlock()

			for key, reloader := range reloaders {
				if reloader.Reload() {
					reloader.BlockReload()
					go func(ctx context.Context, key string, reloader reloadConfig) {
						if err := l.backgroundLoad(ctx, key); err == nil {
							reloader.SetLoadTime(time.Now())
						}
					}(loadCtx, key, reloader)
				}
			}
		case <-l.chClose:
			cancelAll()

			return
		}
	}
}

func valueFrom(returnVal any) (reflect.Value, error) {
	val := reflect.ValueOf(returnVal)

	if val.Kind() != reflect.Pointer {
		return reflect.Value{}, errors.New("value reference must be a pointer type")
	}

	if val.IsZero() {
		val = reflect.New(val.Type())
	}

	return reflect.Indirect(val), nil
}

type loadConfig struct {
	valueType reflect.Type
	parser    Parser
	loader    Loader
	cancel    context.CancelFunc
}

type language struct {
	lang string
}

func (c *language) SetLanguage(lang string) {
	c.lang = lang
}

type reloadConfigSetter func(reloadConfig)

func setReloadConfig(key string, configs map[string]reloadConfig) reloadConfigSetter {
	return func(conf reloadConfig) {
		configs[key] = conf
	}
}

type reloadConfig interface {
	Reload() bool
	BlockReload()
	SetLoadTime(time.Time)
}

type intervalReloadConfig struct {
	mu       sync.RWMutex
	block    bool
	lastLoad time.Time
	interval time.Duration
}

func (c *intervalReloadConfig) Reload() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return time.Since(c.lastLoad) > c.interval && !c.block
}

func (c *intervalReloadConfig) BlockReload() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.block = true
}

func (c *intervalReloadConfig) SetLoadTime(last time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.lastLoad = last
	c.block = false
}

func WithReloadInterval(interval time.Duration) opts.Option {
	return func(val any) {
		if setter, ok := val.(reloadConfigSetter); ok {
			setter(&intervalReloadConfig{
				interval: interval,
			})
		}
	}
}

type loadAfterConfig struct {
	mu      sync.RWMutex
	block   bool
	delay   time.Duration
	created time.Time
	loaded  bool
}

func (c *loadAfterConfig) Reload() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return time.Since(c.created) >= c.delay && !c.loaded && !c.block
}

func (c *loadAfterConfig) BlockReload() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.block = true
}

func (c *loadAfterConfig) SetLoadTime(_ time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.loaded = true
	c.block = false
}

func WithLoadOnceAfter(interval time.Duration) opts.Option {
	return func(val any) {
		if setter, ok := val.(reloadConfigSetter); ok {
			setter(&loadAfterConfig{
				delay:   interval,
				created: time.Now(),
			})
		}
	}
}

type noReloadConfig struct{}

func (c noReloadConfig) Reload() bool {
	// loader is a noop
	return false
}

func (c noReloadConfig) BlockReload() {}

func (c noReloadConfig) SetLoadTime(_ time.Time) {}

func WithNoReload() opts.Option {
	return func(val any) {
		if setter, ok := val.(reloadConfigSetter); ok {
			setter(noReloadConfig{})
		}
	}
}

func makeCacheKey(tp reflect.Type, lang string) string {
	return fmt.Sprintf("%s%+v", lang, tp)
}

type dataCache struct {
	mu       sync.RWMutex
	cache    map[string]reflect.Value
	lastLoad map[string]time.Time
}

func newDataCache() *dataCache {
	return &dataCache{
		cache:    make(map[string]reflect.Value),
		lastLoad: make(map[string]time.Time),
	}
}

func (c *dataCache) get(key string) (reflect.Value, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	value, ok := c.cache[key]

	return value, ok
}

func (c *dataCache) set(key string, value reflect.Value) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = value
	c.lastLoad[key] = time.Now()
}
