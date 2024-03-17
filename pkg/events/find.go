package events

import (
	"encoding/json"
	"fmt"

	deadenz "github.com/ciphermountain/deadenz/pkg"
)

func NewFindEvent(item deadenz.Item) Event {
	return FindEvent{Item: item}
}

type FindEvent struct {
	Item deadenz.Item
}

func (e FindEvent) String() string {
	return fmt.Sprintf("you find %s", e.Item.Name)
}

func (e FindEvent) MarshalJSON() ([]byte, error) {
	type event struct {
		Type string       `json:"type"`
		Item deadenz.Item `json:"item"`
	}

	formatted := event{
		Type: string(EventTypeFind),
		Item: e.Item,
	}

	return json.Marshal(formatted)
}

func (e *FindEvent) UnmarshalJSON(data []byte) error {
	type event struct {
		Item deadenz.Item `json:"item"`
	}

	var formatted event

	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	*e = FindEvent{
		Item: formatted.Item,
	}

	return nil
}
