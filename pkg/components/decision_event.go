package components

import (
	"encoding/json"
	"fmt"
)

// DecisionEvent is intended to be an action on an item found and is not expected to result in a mutation on the
// character.
type ItemDecisionEvent struct {
	value         string
	addToBackpack bool
}

func NewItemDecisionEvent(message string) ItemDecisionEvent {
	return ItemDecisionEvent{
		value: message,
	}
}

// String returns the event in string form.
func (e ItemDecisionEvent) String() string {
	return e.value
}

func (e ItemDecisionEvent) AddToBackpack() bool {
	return e.addToBackpack
}

func (e ItemDecisionEvent) MarshalJSON() ([]byte, error) {
	formatted := jsonItemDecisionEvent{
		Type:    string(EventTypeItemDecision),
		Message: e.value,
	}

	return json.Marshal(formatted)
}

func (e *ItemDecisionEvent) UnmarshalJSON(data []byte) error {
	var formatted jsonItemDecisionEvent
	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	if formatted.Type != string(EventTypeItemDecision) {
		return fmt.Errorf("%w: %s; expected %s", ErrInvalidEventType, formatted.Type, EventTypeItemDecision)
	}

	*e = ItemDecisionEvent{
		value:         formatted.Message,
		addToBackpack: formatted.AddToBackpack,
	}

	return nil
}

type jsonItemDecisionEvent struct {
	Type          string `json:"type"`
	Message       string `json:"message"`
	AddToBackpack bool   `json:"addToBackpack"`
}
