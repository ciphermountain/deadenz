package events

import (
	"encoding/json"
	"fmt"

	"github.com/ciphermountain/deadenz/pkg/components"
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

func (e EncounterEvent) MarshalJSON() ([]byte, error) {
	type event struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}

	formatted := event{
		Type:    string(components.EventTypeEncounter),
		Message: e.value,
	}

	return json.Marshal(formatted)
}

func (e *EncounterEvent) UnmarshalJSON(data []byte) error {
	type event struct {
		Message string `json:"message"`
	}

	var formatted event

	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	*e = EncounterEvent{
		value: formatted.Message,
	}

	return nil
}
