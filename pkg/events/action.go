package events

import (
	"encoding/json"

	"github.com/ciphermountain/deadenz/pkg/components"
)

func LoadActionEvents(b []byte) ([]ActionEvent, error) {
	var loaded []ActionEvent

	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, err
	}

	return loaded, nil
}

// ActionEvent is intended to be something a character does. This can have effects on the character.
type ActionEvent struct {
	value string
}

func (e ActionEvent) String() string {
	return e.value
}

func (e ActionEvent) MarshalJSON() ([]byte, error) {
	type action struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}

	formatted := action{
		Type:    string(components.EventTypeAction),
		Message: e.value,
	}

	return json.Marshal(formatted)
}

func (e *ActionEvent) UnmarshalJSON(data []byte) error {
	type action struct {
		Message string `json:"message"`
	}

	var formatted action

	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	*e = ActionEvent{
		value: formatted.Message,
	}

	return nil
}
