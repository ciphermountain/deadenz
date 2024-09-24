package components

import (
	"encoding/json"
	"errors"
	"fmt"
)

type DieMutationEvent struct {
	value string
}

func NewDieMutationEvent(value string) DieMutationEvent {
	return DieMutationEvent{
		value: value,
	}
}

func (e DieMutationEvent) String() string {
	return e.value
}

func (e DieMutationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonMutationEvent{
		Type:    string(EventTypeMutation),
		Message: e.value,
		IsDeath: true,
	})
}

func (e *DieMutationEvent) UnmarshalJSON(data []byte) error {
	var formatted jsonMutationEvent
	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	if !formatted.IsDeath {
		return errors.New("not a death event")
	}

	*e = DieMutationEvent{
		value: formatted.Message,
	}

	return nil
}

type DieMutationEventWithCharacter struct {
	Character CharacterType
	Death     DieMutationEvent
}

func NewDieMutationEventWithCharacter(character Character, evt DieMutationEvent) DieMutationEventWithCharacter {
	return DieMutationEventWithCharacter{
		Character: character.Type,
		Death:     evt,
	}
}

func (e DieMutationEventWithCharacter) String() string {
	return e.Death.value
}

func (e DieMutationEventWithCharacter) MarshalJSON() ([]byte, error) {
	formatted := jsonMutationEvent{
		Type:      string(EventTypeMutation),
		Message:   e.Death.value,
		IsDeath:   true,
		Character: (*uint64)(&e.Character),
	}

	return json.Marshal(formatted)
}

func (e *DieMutationEventWithCharacter) UnmarshalJSON(data []byte) error {
	var formatted jsonMutationEvent

	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	if !formatted.IsDeath {
		return errors.New("not a death event")
	}

	if formatted.Character == nil {
		return errors.New("death event has no character")
	}

	*e = DieMutationEventWithCharacter{
		Character: CharacterType(*formatted.Character),
		Death: DieMutationEvent{
			value: formatted.Message,
		},
	}

	return nil
}

type LiveMutationEvent struct {
	value string
}

func NewLiveMutationEvent(value string) LiveMutationEvent {
	return LiveMutationEvent{
		value: value,
	}
}

func (e LiveMutationEvent) String() string {
	return e.value
}

func (e LiveMutationEvent) MarshalJSON() ([]byte, error) {
	formatted := jsonMutationEvent{
		Type:    string(EventTypeMutation),
		Message: e.value,
		IsDeath: false,
	}

	return json.Marshal(formatted)
}

func (e *LiveMutationEvent) UnmarshalJSON(data []byte) error {
	var formatted jsonMutationEvent
	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	if formatted.Type != string(EventTypeMutation) {
		return fmt.Errorf("%w: %s; expected %s", ErrInvalidEventType, formatted.Type, EventTypeMutation)
	}

	if formatted.IsDeath {
		return ErrNotLiveEvent
	}

	*e = LiveMutationEvent{
		value: formatted.Message,
	}

	return nil
}

type jsonMutationEvent struct {
	Type      string  `json:"type"`
	Message   string  `json:"message"`
	IsDeath   bool    `json:"isDeath"`
	Trap      bool    `json:"trap"`
	Character *uint64 `json:"character_type,omitempty"`
}

type TripTrapMutationEvent struct {
	msg string
}

func NewTripTrapMutationEvent(msg string) Event {
	return Event{typed: TripTrapMutationEvent{
		msg: msg,
	}}
}

func (e TripTrapMutationEvent) String() string {
	return e.msg
}

func (e TripTrapMutationEvent) MarshalJSON() ([]byte, error) {
	formatted := jsonMutationEvent{
		Type:    string(EventTypeMutation),
		Message: e.msg,
		IsDeath: true,
		Trap:    true,
	}

	return json.Marshal(formatted)
}

func (e *TripTrapMutationEvent) UnmarshalJSON(data []byte) error {
	var formatted jsonMutationEvent
	if err := json.Unmarshal(data, &formatted); err != nil {
		return err
	}

	if !formatted.IsDeath {
		return errors.New("not a death event")
	}

	if !formatted.Trap {
		return errors.New("not a trap mutation event")
	}

	*e = TripTrapMutationEvent{
		msg: formatted.Message,
	}

	return nil
}
