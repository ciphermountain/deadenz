package events

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ciphermountain/deadenz/pkg/components"
)

func NewCharacterSpawnEvent(character components.Character) components.Event {
	return &CharacterSpawnEvent{
		character: character}
}

type CharacterSpawnEvent struct {
	character components.Character
}

func (e CharacterSpawnEvent) String() string {
	return fmt.Sprintf("you spawned in as a %s", e.character.Name)
}

func (e CharacterSpawnEvent) Type() components.CharacterType {
	return e.character.Type
}

func (e CharacterSpawnEvent) Name() string {
	return e.character.Name
}

func (e CharacterSpawnEvent) MarshalJSON() ([]byte, error) {
	bts, err := json.Marshal(e.character)
	if err != nil {
		return nil, err
	}

	msg := json.RawMessage(bts)

	return json.Marshal(typer{
		Type: "spawnin-event",
		Data: &msg,
	})
}

func (e *CharacterSpawnEvent) UnmarshalJSON(bts []byte) error {
	var base typer
	if err := json.Unmarshal(bts, &base); err != nil {
		return err
	}

	if base.Data == nil {
		return errors.New("invalid spawnin event")
	}

	var character components.Character
	if err := json.Unmarshal(*base.Data, &character); err != nil {
		return err
	}

	*e = CharacterSpawnEvent{character: character}

	return nil
}

type typer struct {
	Type string           `json:"type"`
	Data *json.RawMessage `json:"data"`
}

func NewEarnedXPEvent(xp uint) components.Event {
	return &EarnedXPEvent{xp: xp}
}

type EarnedXPEvent struct {
	xp uint
}

func (e EarnedXPEvent) String() string {
	return fmt.Sprintf("you earned %d xp", e.xp)
}

func NewEarnedTokenEvent(xp uint) components.Event {
	return &EarnedTokenEvent{xp: xp}
}

type EarnedTokenEvent struct {
	xp uint
}

func (e EarnedTokenEvent) String() string {
	return fmt.Sprintf("you earned %d tokens", e.xp)
}
