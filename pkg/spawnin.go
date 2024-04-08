package deadenz

import (
	"context"
	"errors"

	"github.com/ciphermountain/deadenz/internal/util"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/events"
)

var (
	ErrAlreadySpawnedIn = errors.New("character already exists; cannot spawn in")
)

type Loader interface {
	// Load is expected to load any value immediately
	Load(any) error
	// LoadCtx provides a context that can establish a deadline
	LoadCtx(context.Context, any) error
}

// Spawn assigns a new character to an existing profile and modifies xp, backpack,
// and stats. Events emitted include spawn event and earned xp event. Will return
// an already spawned error if profile has an active character.
func Spawn(profile components.Profile, loader Loader) (components.Profile, []components.Event, error) {
	// short circuit if the user has an active character
	if profile.Active != nil {
		return profile, nil, ErrAlreadySpawnedIn
	}

	var characters []components.Character
	if err := loader.Load(&characters); err != nil {
		return profile, nil, err
	}

	char := characters[util.Random(0, int64(len(characters)-1))]

	profile.XP = profile.XP + uint(char.Multiplier)
	profile.Active = &char
	// TODO: register in multiverse

	evts := []components.Event{
		events.NewCharacterSpawnEvent(char),
		events.NewEarnedXPEvent(uint(char.Multiplier)),
	}

	return profile, evts, nil
}
