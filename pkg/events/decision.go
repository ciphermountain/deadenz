package events

import (
	"encoding/json"
)

func LoadItemDecisions(b []byte) ([]ItemDecisionEvent, error) {
	type decision struct {
		Message       string `json:"message"`
		AddToBackpack bool   `json:"addToBackpack"`
	}

	var loaded []decision

	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, err
	}

	evts := []ItemDecisionEvent{}

	for _, l := range loaded {
		evts = append(evts, ItemDecisionEvent{
			value:         l.Message,
			addToBackpack: l.AddToBackpack,
		})
	}

	return evts, nil
}

// DecisionEvent is intended to be an action on an item found and is not expected to result in a mutation on the
// character.
type ItemDecisionEvent struct {
	value         string
	addToBackpack bool
}

// String returns the event in string form.
func (e ItemDecisionEvent) String() string {
	return e.value
}

func (e ItemDecisionEvent) AddToBackpack() bool {
	return e.addToBackpack
}
