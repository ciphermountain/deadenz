package events

import (
	"encoding/json"

	"github.com/ciphermountain/deadenz/pkg/components"
)

func LoadItemDecisions(b []byte) ([]ItemDecisionEvent, error) {
	var loaded []ItemDecisionEvent

	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, err
	}

	return loaded, nil
}

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
	type event struct {
		Type          string `json:"type"`
		Message       string `json:"message"`
		AddToBackpack bool   `json:"addToBackpack"`
	}

	formatted := event{
		Type:    string(components.EventTypeItemDecision),
		Message: e.value,
	}

	return json.Marshal(formatted)
}

func (e *ItemDecisionEvent) UnmarshalJSON(data []byte) error {
	type event struct {
		Message       string `json:"message"`
		AddToBackpack bool   `json:"addToBackpack"`
	}

	var formatted event

	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	*e = ItemDecisionEvent{
		value: formatted.Message,
	}

	return nil
}
