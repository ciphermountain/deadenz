package components

import (
	"encoding/json"
	"fmt"
)

type FindEvent struct {
	Item Item
}

func NewFindEvent(item Item) FindEvent {
	return FindEvent{Item: item}
}

func (e FindEvent) String() string {
	return fmt.Sprintf("you find %s", e.Item.Name) // TODO: breaks multi-language support
}

func (e FindEvent) MarshalJSON() ([]byte, error) {
	formatted := jsonFindEvent{
		Type: string(EventTypeFind),
		Item: e.Item,
	}

	return json.Marshal(formatted)
}

func (e *FindEvent) UnmarshalJSON(data []byte) error {
	var formatted jsonFindEvent
	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	if formatted.Type != string(EventTypeFind) {
		return fmt.Errorf("%w: %s; expected %s", ErrInvalidEventType, formatted.Type, EventTypeFind)
	}

	*e = FindEvent{
		Item: formatted.Item,
	}

	return nil
}

type jsonFindEvent struct {
	Type string `json:"type"`
	Item Item   `json:"item"`
}
