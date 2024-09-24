package components

import (
	"encoding/json"
	"fmt"
)

type CharacterSpawnEvent struct {
	character Character
}

func NewCharacterSpawnEvent(character Character) *CharacterSpawnEvent {
	return &CharacterSpawnEvent{
		character: character}
}

func (e CharacterSpawnEvent) String() string {
	return fmt.Sprintf("you spawned in as a %s", e.character.Name) // TODO: breaks multi-language support
}

func (e CharacterSpawnEvent) Type() CharacterType {
	return e.character.Type
}

func (e CharacterSpawnEvent) Name() string {
	return e.character.Name
}

func (e CharacterSpawnEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonCharacterSpawnEvent{
		Type:      "spawnin-event",
		Character: e.character,
	})
}

func (e *CharacterSpawnEvent) UnmarshalJSON(bts []byte) error {
	var base jsonCharacterSpawnEvent
	if err := json.Unmarshal(bts, &base); err != nil {
		return err
	}

	if base.Type != "spawnin-event" {
		return fmt.Errorf("%w: %s; expected spawnin-event", ErrInvalidEventType, base.Type)
	}

	*e = CharacterSpawnEvent{character: base.Character}

	return nil
}

type jsonCharacterSpawnEvent struct {
	Type      string    `json:"type"`
	Character Character `json:"character"`
}

type EarnedXPEvent struct {
	xp uint
}

func NewEarnedTokenEvent(xp uint) *EarnedTokenEvent {
	return &EarnedTokenEvent{xp: xp}
}

func (e EarnedXPEvent) String() string {
	return fmt.Sprintf("you earned %d xp", e.xp) // TODO: breaks multi-language support
}

type EarnedTokenEvent struct {
	xp uint
}

func NewEarnedXPEvent(xp uint) *EarnedXPEvent {
	return &EarnedXPEvent{xp: xp}
}

func (e EarnedTokenEvent) String() string {
	return fmt.Sprintf("you earned %d tokens", e.xp) // TODO: breaks multi-language support
}
