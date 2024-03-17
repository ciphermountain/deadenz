package multiverse

import (
	"context"
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"

	"github.com/ciphermountain/deadenz/pkg/multiverse/service"
)

var (
	ErrUnimplemented = errors.New("unimplemented")
)

var _ service.MultiverseServer = &MultiverseServer{}

type MultiverseServer struct {
	service.UnimplementedMultiverseServer
	subscribers map[string]chan *service.Event
	mu          sync.RWMutex
}

func NewMultiverseServer() *MultiverseServer {
	return &MultiverseServer{
		subscribers: make(map[string]chan *service.Event),
	}
}

func (s *MultiverseServer) PublishEvent(ctx context.Context, event *service.Event) (*service.Response, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, chEvents := range s.subscribers {
		go func(ch chan *service.Event, evt *service.Event) {
			select {
			case ch <- evt:
				return
			case <-ctx.Done():
				return
			}
		}(chEvents, event)
	}

	return nil, ErrUnimplemented
}

func (s *MultiverseServer) Events(_ *service.Filter, stream service.Multiverse_EventsServer) error {
	name := uuid.NewV4()
	chEvents := make(chan *service.Event, 100)

	s.mu.Lock()
	s.subscribers[name.String()] = chEvents
	s.mu.Unlock()

	// immediately put an event on the channel with the subscription name

	for {
		select {
		case <-stream.Context().Done():
			s.mu.Lock()
			delete(s.subscribers, name.String())
			s.mu.Unlock()

			return nil
		case evt := <-chEvents:
			if err := stream.Send(evt); err != nil {
				return err
			}
		}
	}
}

func MultiverseDeathMiddleware() {
	// spawnin events are registered per player
	// death events are registered
}
