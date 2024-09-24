package components

import (
	"encoding/json"
	"fmt"
)

// ActionEvent is intended to be something a character does. This can have effects on the character.
type ActionEvent struct {
	value string
}

func NewActionEvent(value string) ActionEvent {
	return ActionEvent{value: value}
}

func (e ActionEvent) String() string {
	return e.value
}

func (e ActionEvent) MarshalJSON() ([]byte, error) {
	formatted := jsonActionEvent{
		Type:    string(EventTypeAction),
		Message: e.value,
	}

	return json.Marshal(formatted)
}

func (e *ActionEvent) UnmarshalJSON(data []byte) error {
	var formatted jsonActionEvent
	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	if formatted.Type != string(EventTypeAction) {
		return fmt.Errorf("%w: %s; expected %s", ErrInvalidEventType, formatted.Type, EventTypeAction)
	}

	*e = ActionEvent{
		value: formatted.Message,
	}

	return nil
}

type jsonActionEvent struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
