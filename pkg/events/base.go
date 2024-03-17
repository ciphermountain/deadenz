package events

import (
	"encoding/json"
	"errors"
)

type Event interface {
	String() string
}

type EventType string

const (
	EventTypeAction       EventType = "action"
	EventTypeItemDecision EventType = "item_decision"
	EventTypeEncounter    EventType = "encounter"
	EventTypeFind         EventType = "find"
	EventTypeMutation     EventType = "mutation"
	EventTypeSpawnin      EventType = "spawnin"
)

func DecodeJSONEvent(data []byte) (Event, error) {
	type typer struct {
		Type string `json:"type"`
	}

	var onlyType typer

	if err := json.Unmarshal(data, &onlyType); err != nil {
		return nil, err
	}

	switch EventType(onlyType.Type) {
	case EventTypeAction:
		var event ActionEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}

		return event, nil
	case EventTypeItemDecision:
		var event ItemDecisionEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}

		return event, nil
	case EventTypeEncounter:
		var event EncounterEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}

		return event, nil
	case EventTypeFind:
		var event FindEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}

		return event, nil
	case EventTypeMutation:
		return UnmarshalMutationEvent(data)
	default:
		return nil, errors.New("unknown event type")
	}
}
