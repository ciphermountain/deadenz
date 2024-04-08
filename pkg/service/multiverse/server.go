package multiverse

import (
	"context"
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"

	proto "github.com/ciphermountain/deadenz/pkg/proto/multiverse"
)

var (
	ErrUnimplemented = errors.New("unimplemented")
)

var _ proto.MultiverseServer = &MultiverseServer{}

type MultiverseServer struct {
	proto.UnimplementedMultiverseServer
	subscribers map[string]chan *proto.Event
	mu          sync.RWMutex
}

func NewMultiverseServer() *MultiverseServer {
	return &MultiverseServer{
		subscribers: make(map[string]chan *proto.Event),
	}
}

func (s *MultiverseServer) PublishEvent(ctx context.Context, event *proto.Event) (*proto.Response, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, chEvents := range s.subscribers {
		go func(ch chan *proto.Event, evt *proto.Event) {
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

func (s *MultiverseServer) Events(_ *proto.Filter, stream proto.Multiverse_EventsServer) error {
	name := uuid.NewV4()
	chEvents := make(chan *proto.Event, 100)

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
