package events

import (
	"encoding/json"
	"fmt"

	"github.com/ciphermountain/deadenz/pkg/components"
)

func NewFindEvent(item components.Item) components.Event {
	return FindEvent{Item: item}
}

type FindEvent struct {
	Item components.Item
}

func (e FindEvent) String() string {
	return fmt.Sprintf("you find %s", e.Item.Name) // TODO: breaks multi-language support
}

func (e FindEvent) MarshalJSON() ([]byte, error) {
	type event struct {
		Type string          `json:"type"`
		Item components.Item `json:"item"`
	}

	formatted := event{
		Type: string(components.EventTypeFind),
		Item: e.Item,
	}

	return json.Marshal(formatted)
}

func (e *FindEvent) UnmarshalJSON(data []byte) error {
	type event struct {
		Item components.Item `json:"item"`
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
