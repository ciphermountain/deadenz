package components

import (
	"encoding/json"
	"errors"
	"fmt"
)

type EventType string

const (
	EventTypeAction       EventType = "action"
	EventTypeItemDecision EventType = "item_decision"
	EventTypeEncounter    EventType = "encounter"
	EventTypeFind         EventType = "find"
	EventTypeMutation     EventType = "mutation"
	EventTypeSpawnin      EventType = "spawnin"
)

type Event struct {
	typed any
}

func NewEvent(evt any) Event {
	return Event{typed: evt}
}

func (e Event) String() string {
	evt, ok := e.typed.(fmt.Stringer)
	if !ok {
		panic(fmt.Sprintf("%T event should implement type fmt.Stringer", evt))
	}

	return evt.String()
}

func (e Event) Typed() any {
	return e.typed
}

func (e Event) MarshalJSON() ([]byte, error) {
	return nil, errors.New("unimplemented")
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type typer struct {
		Type string `json:"type"`
	}

	var onlyType typer

	if err := json.Unmarshal(data, &onlyType); err != nil {
		return err
	}

	switch EventType(onlyType.Type) {
	case EventTypeAction:
		var event ActionEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return err
		}

		*e = Event{typed: event}
	case EventTypeItemDecision:
		var event ItemDecisionEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return err
		}

		*e = Event{typed: event}
	case EventTypeEncounter:
		var event EncounterEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return err
		}

		*e = Event{typed: event}
	case EventTypeFind:
		var event FindEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return err
		}

		*e = Event{typed: event}
	case EventTypeMutation:
		event, err := unmarshalMutationEvent(data)
		if err != nil {
			return err
		}

		*e = Event{typed: event}
	case EventTypeSpawnin:
		var event CharacterSpawnEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return err
		}

		*e = Event{typed: event}
	default:
		return errors.New("unknown event type")
	}

	return nil
}

func unmarshalMutationEvent(data []byte) (any, error) {
	type action struct {
		Message   string  `json:"message"`
		IsDeath   bool    `json:"isDeath"`
		Character *uint64 `json:"character_type"`
		Trap      bool    `json:"trap"`
	}

	var loaded action

	if err := json.Unmarshal(data, &loaded); err != nil {
		return nil, err
	}

	if loaded.IsDeath {
		if loaded.Trap {
			return NewTripTrapMutationEvent(loaded.Message), nil
		}

		evt := NewDieMutationEvent(loaded.Message)

		if loaded.Character != nil {
			return NewDieMutationEventWithCharacter(
				Character{Type: CharacterType(*loaded.Character)},
				evt,
			), nil
		}

		return evt, nil
	}

	return NewLiveMutationEvent(loaded.Message), nil
}
