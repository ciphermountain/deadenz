package components

import (
	"encoding/json"
	"fmt"
)

type EncounterEvent struct {
	value string
}

func NewEncounterEvent(value string) EncounterEvent {
	return EncounterEvent{value: value}
}

func (e EncounterEvent) String() string {
	return e.value
}

func (e EncounterEvent) MarshalJSON() ([]byte, error) {
	formatted := jsonEncounterEvent{
		Type:    string(EventTypeEncounter),
		Message: e.value,
	}

	return json.Marshal(formatted)
}

func (e *EncounterEvent) UnmarshalJSON(data []byte) error {
	var formatted jsonEncounterEvent
	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	if formatted.Type != string(EventTypeEncounter) {
		return fmt.Errorf("%w: %s; expected %s", ErrInvalidEventType, formatted.Type, EventTypeEncounter)
	}

	*e = EncounterEvent{
		value: formatted.Message,
	}

	return nil
}

type jsonEncounterEvent struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
