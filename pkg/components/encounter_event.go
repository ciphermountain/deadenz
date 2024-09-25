package components

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ciphermountain/deadenz/pkg/opts"
)

type EncounterEvent struct {
	value string
	lang  LanguagePack
}

func NewEncounterEvent(value string, opts ...opts.Option) EncounterEvent {
	lang := &language{}

	for _, opt := range opts {
		opt(lang)
	}

	return EncounterEvent{value: value, lang: lang.lang}
}

func (e EncounterEvent) WithOpts(opts ...opts.Option) EncounterEvent {
	return NewEncounterEvent(e.value, opts...)
}

func (e EncounterEvent) String() string {
	return strings.ReplaceAll(e.lang.EncounterPattern, "{{encounter}}", e.value)
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
