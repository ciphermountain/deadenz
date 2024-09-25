package data

import (
	"context"
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
	mu sync.RWMutex
	// data    map[reflect.Type]any
	configs map[string]*loadConfig

	starter sync.Once
	closer  sync.Once
	running atomic.Bool
	chClose chan struct{}
}

func NewDataLoader() *DataLoader {
	return &DataLoader{
		// data:    make(map[reflect.Type]any),
		configs: make(map[string]*loadConfig),
		chClose: make(chan struct{}, 1),
	}
}

func (l *DataLoader) SetLoader(typeKey reflect.Type, loader Loader, parser Parser, options ...opts.Option) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	language := &language{lang: ""}
	config := &loadConfig{
		parser: parser,
		loader: loader,
	}

	for _, opt := range options {
		opt(language)
		opt(config)
	}

	key := fmt.Sprintf("%s%+v", language.lang, typeKey)

	l.configs[key] = config

	return nil
}

func (l *DataLoader) Load(value any, options ...opts.Option) error {
	return l.LoadCtx(context.Background(), value, options...)
}

func (l *DataLoader) LoadCtx(ctx context.Context, value any, options ...opts.Option) error {
	l.mu.RLock()
	defer l.mu.RUnlock()

	val := reflect.ValueOf(value)
	tp := reflect.Indirect(val).Type()
	language := &language{lang: ""}

	for _, opt := range options {
		opt(language)
	}

	key := fmt.Sprintf("%s%+v", language.lang, tp)
	// log.Println(key)

	/*
		if data, exists := l.data[tp]; exists {
			reflect.Indirect(val).Set(reflect.Indirect(reflect.ValueOf(data)))

			return nil
		}
	*/

	config, exists := l.configs[key]
	if !exists || config.loader == nil {
		return fmt.Errorf("loader does not exist for %+v", tp)
	}

	if config.parser == nil {
		return fmt.Errorf("parser does not exist for %+v", tp)
	}

	bts, err := config.loader.Data(ctx)
	if err != nil {
		return err
	}

	if err := config.parser(bts, value); err != nil {
		return err
	}

	// l.data[tp] = value

	return nil
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
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			// search configs and run loaders
		case <-l.chClose:
			return
		}
	}
}

type loadConfig struct {
	interval time.Duration
	reload   bool
	delay    bool
	parser   Parser
	loader   Loader
}

type language struct {
	lang string
}

func (c *language) SetLanguage(lang string) {
	c.lang = lang
}

func WithReloadInterval(interval time.Duration) opts.Option {
	return func(val any) {
		if conf, ok := val.(*loadConfig); ok {
			conf.reload = true
			conf.interval = interval
		}
	}
}

func WithLoadOnceAfter(interval time.Duration) opts.Option {
	return func(val any) {
		if conf, ok := val.(*loadConfig); ok {
			conf.delay = true
			conf.interval = interval
		}
	}
}

func WithNoReload() opts.Option {
	return func(_ any) {}
}
