package util

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

type Parser func([]byte, any) error

type Loader interface {
	Data(context.Context) ([]byte, error)
}

type LoaderOpt func(conf *loadConfig)

type DataLoader struct {
	mu sync.RWMutex
	// data    map[reflect.Type]any
	configs map[reflect.Type]*loadConfig

	starter sync.Once
	closer  sync.Once
	running atomic.Bool
	chClose chan struct{}
}

func NewDataLoader() *DataLoader {
	return &DataLoader{
		// data:    make(map[reflect.Type]any),
		configs: make(map[reflect.Type]*loadConfig),
		chClose: make(chan struct{}, 1),
	}
}

func (l *DataLoader) SetLoader(key reflect.Type, loader Loader, parser Parser, opts ...LoaderOpt) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	config := loadConfig{
		parser: parser,
		loader: loader,
	}

	for _, opt := range opts {
		opt(&config)
	}

	l.configs[key] = &config

	return nil
}

func (l *DataLoader) Load(value any) error {
	return l.LoadCtx(context.Background(), value)
}

func (l *DataLoader) LoadCtx(ctx context.Context, value any) error {
	l.mu.RLock()
	defer l.mu.RUnlock()

	val := reflect.ValueOf(value)
	tp := reflect.Indirect(val).Type()

	/*
		if data, exists := l.data[tp]; exists {
			reflect.Indirect(val).Set(reflect.Indirect(reflect.ValueOf(data)))

			return nil
		}
	*/

	config, exists := l.configs[tp]
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

func WithReloadInterval(interval time.Duration) LoaderOpt {
	return func(conf *loadConfig) {
		conf.reload = true
		conf.interval = interval
	}
}

func WithLoadOnceAfter(interval time.Duration) LoaderOpt {
	return func(conf *loadConfig) {
		conf.delay = true
		conf.interval = interval
	}
}

func WithNoReload() LoaderOpt {
	return func(conf *loadConfig) {}
}
