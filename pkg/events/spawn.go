package events

import (
	"fmt"

	deadenz "github.com/ciphermountain/deadenz/pkg"
)

func NewCharacterSpawnEvent(character deadenz.Character) Event {
	return &CharacterSpawnEvent{
		character: character}
}

type CharacterSpawnEvent struct {
	character deadenz.Character
}

func (e CharacterSpawnEvent) String() string {
	return fmt.Sprintf("you spawned in as a %s", e.character.Name)
}

func NewEarnedXPEvent(xp uint) Event {
	return &EarnedXPEvent{xp: xp}
}

type EarnedXPEvent struct {
	xp uint
}

func (e EarnedXPEvent) String() string {
	return fmt.Sprintf("you earned %d xp", e.xp)
}

func NewEarnedTokenEvent(xp uint) Event {
	return &EarnedTokenEvent{xp: xp}
}

type EarnedTokenEvent struct {
	xp uint
}

func (e EarnedTokenEvent) String() string {
	return fmt.Sprintf("you earned %d tokens", e.xp)
}
