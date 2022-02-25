package core

import "fmt"

type Event interface {
	String() string
}

func NewCharacterSpawnedInEvent(character Character) Event {
	return &CharacterSpawnedInEvent{
		character: character}
}

type CharacterSpawnedInEvent struct {
	character Character
}

func (e *CharacterSpawnedInEvent) String() string {
	return fmt.Sprintf("you spawned in as a %s", e.character.Name())
}
