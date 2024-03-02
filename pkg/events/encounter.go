package events

import (
	"encoding/json"
	"fmt"
)

func LoadEncounterEvents(b []byte) ([]EncounterEvent, error) {
	type event struct {
		Message string `json:"message"`
	}

	var loaded []event

	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, err
	}

	evts := []EncounterEvent{}

	for _, l := range loaded {
		evts = append(evts, EncounterEvent{
			value: l.Message,
		})
	}

	return evts, nil
}

type EncounterEvent struct {
	value string
}

func (e EncounterEvent) String() string {
	return fmt.Sprintf("you encounter %s", e.value)
}
