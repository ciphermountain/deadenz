package multiverse

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/events"
	"github.com/ciphermountain/deadenz/pkg/parse"
	proto "github.com/ciphermountain/deadenz/pkg/proto/multiverse"
)

var (
	ErrUnimplemented = errors.New("unimplemented")
)

var _ proto.MultiverseServer = &MultiverseServer{}

type MultiverseServer struct {
	proto.UnimplementedMultiverseServer
	subscribers      map[string]chan *proto.Event
	characterLookup  map[components.CharacterType][]string
	characterReverse map[string]components.CharacterType
	characterLocks   map[components.CharacterType]*sync.RWMutex
	mu               sync.RWMutex
}

func NewMultiverseServer() *MultiverseServer {
	return &MultiverseServer{
		subscribers:      make(map[string]chan *proto.Event),
		characterLookup:  make(map[components.CharacterType][]string),
		characterReverse: make(map[string]components.CharacterType),
		characterLocks:   make(map[components.CharacterType]*sync.RWMutex),
	}
}

func (s *MultiverseServer) PublishGameEvent(ctx context.Context, event *proto.GameEvent) (*proto.Response, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// unmarshal game event
	evt, err := parse.DecodeJSONEvent(event.Data)
	if err != nil {
		return nil, err
	}

	switch typed := evt.(type) {
	case events.CharacterSpawnEvent:
		s.saveSpawnEvent(typed, event.Uid)
	case events.DieMutationEventWithCharacter:
		s.processDeathEvent(typed)
	}

	return &proto.Response{
		Status: proto.Status_OK,
	}, nil
}

func (s *MultiverseServer) Events(filter *proto.Filter, stream proto.Multiverse_EventsServer) error {
	chEvents := make(chan *proto.Event, 100)

	s.mu.Lock()
	s.subscribers[filter.Uid] = chEvents
	s.mu.Unlock()

	for {
		select {
		case <-stream.Context().Done():
			s.mu.Lock()
			delete(s.subscribers, filter.Uid)
			s.mu.Unlock()

			return nil
		case evt := <-chEvents:
			if err := stream.Send(evt); err != nil {
				return err
			}
		}
	}
}

func (s *MultiverseServer) saveSpawnEvent(evt events.CharacterSpawnEvent, id string) {
	lock := mustGetCharacterLock(evt.Type(), s.characterLocks)

	lock.Lock()
	s.characterLookup[evt.Type()] = mustAddValue(s.characterLookup[evt.Type()], id)
	lock.Unlock()

	s.mu.Lock()
	s.characterReverse[id] = evt.Type()
	s.mu.Unlock()
}

func (s *MultiverseServer) processDeathEvent(evt events.DieMutationEventWithCharacter) {
	character := evt.Character
	lock := mustGetCharacterLock(character, s.characterLocks)

	lock.Lock()
	recipients := s.characterLookup[character]

	delete(s.characterLookup, character)
	lock.Unlock()

	event := &proto.Event{
		Type: &proto.Event_CharacterDeath{
			CharacterDeath: &proto.DeathByCharacterType{
				Type: uint64(character),
			},
		},
	}

	for _, recipient := range recipients {
		s.mu.RLock()
		chEvt, ok := s.subscribers[recipient]
		s.mu.RUnlock()

		if !ok {
			continue
		}

		go sendToChannel(chEvt, event)
	}
}

func mustAddValue(values []string, toAdd string) []string {
	for _, val := range values {
		if val == toAdd {
			return values
		}
	}

	return append(values, toAdd)
}

func mustGetCharacterLock(character components.CharacterType, locks map[components.CharacterType]*sync.RWMutex) *sync.RWMutex {
	lock, ok := locks[character]
	if !ok {
		lock = &sync.RWMutex{}
		locks[character] = lock
	}

	return lock
}

func sendToChannel(chEvts chan *proto.Event, evt *proto.Event) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	select {
	case chEvts <- evt:
		cancel()
	case <-ctx.Done():
		cancel()
	}
}
