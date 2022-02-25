package core

import (
	"errors"
)

var (
	ErrAlreadySpawnedIn = errors.New("character already exists; cannot spawn in")
)

func SpawnIn(profile Profile) (Profile, Event, error) {
	if profile.Active != nil {
		return profile, nil, ErrAlreadySpawnedIn
	}

	character := NewCharacter(MagicianCharacterType)
	event := NewCharacterSpawnedInEvent(character)

	profile.Active = &character

	return profile, event, nil
}
