package components

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ciphermountain/deadenz/pkg/opts"
)

type FindEvent struct {
	Item Item
	lang LanguagePack
}

func NewFindEvent(item Item, opts ...opts.Option) FindEvent {
	lang := &language{}

	for _, opt := range opts {
		opt(lang)
	}

	return FindEvent{Item: item, lang: lang.lang}
}

func (e FindEvent) String() string {
	return strings.ReplaceAll(e.lang.FindPattern, "{{item}}", e.Item.Name)
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
