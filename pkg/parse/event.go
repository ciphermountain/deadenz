package parse

import (
	"encoding/json"
	"errors"

	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/events"
)

func DecodeJSONEvent(data []byte) (components.Event, error) {
	type typer struct {
		Type string `json:"type"`
	}

	var onlyType typer

	if err := json.Unmarshal(data, &onlyType); err != nil {
		return nil, err
	}

	switch components.EventType(onlyType.Type) {
	case components.EventTypeAction:
		var event events.ActionEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}

		return event, nil
	case components.EventTypeItemDecision:
		var event events.ItemDecisionEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}

		return event, nil
	case components.EventTypeEncounter:
		var event events.EncounterEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}

		return event, nil
	case components.EventTypeFind:
		var event events.FindEvent
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}

		return event, nil
	case components.EventTypeMutation:
		return UnmarshalMutationEvent(data)
	default:
		return nil, errors.New("unknown event type")
	}
}

func UnmarshalMutationEvent(data []byte) (components.Event, error) {
	type action struct {
		Message string `json:"message"`
		IsDeath bool   `json:"isDeath"`
	}

	var loaded action

	if err := json.Unmarshal(data, &loaded); err != nil {
		return nil, err
	}

	if loaded.IsDeath {
		return events.NewDieMutationEvent(loaded.Message), nil
	}

	return events.NewLiveMutationEvent(loaded.Message), nil
}
