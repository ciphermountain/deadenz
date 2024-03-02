package events

import (
	"encoding/json"
)

func LoadActionEvents(b []byte) ([]ActionEvent, error) {
	type action struct {
		Message string `json:"message"`
	}

	var loaded []action

	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, err
	}

	evts := []ActionEvent{}

	for _, l := range loaded {
		evts = append(evts, ActionEvent{
			value: l.Message,
		})
	}

	return evts, nil
}

// ActionEvent is intended to be something a character does. This can have effects on the character.
type ActionEvent struct {
	value string
}

func (e ActionEvent) String() string {
	return e.value
}
