package deadenz

import (
	"context"
	"errors"
	"fmt"

	"github.com/ciphermountain/deadenz/internal/util"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/opts"
)

var (
	ErrAlreadySpawnedIn = errors.New("character already exists; cannot spawn in")
)

type Loader interface {
	// Load is expected to load any value immediately
	Load(any, ...opts.Option) error
	// LoadCtx provides a context that can establish a deadline
	LoadCtx(context.Context, any, ...opts.Option) error
}

type Config struct {
	ItemFindRate     float64
	TrapTripRate     float64
	DeathRate        float64
	WalkLimitPerHour uint16
	Language         string
}

// Spawn assigns a new character to an existing profile and modifies xp, backpack,
// and stats. Events emitted include spawn event and earned xp event. Will return
// an already spawned error if profile has an active character.
func Spawn(
	profile *components.Profile,
	loader Loader,
	_ Config,
	options ...opts.Option,
) (*components.Profile, []components.Event, error) {
	// short circuit if the user has an active character
	if profile.Active != nil {
		return profile, nil, ErrAlreadySpawnedIn
	}

	var characters []components.Character
	if err := loader.Load(&characters, options...); err != nil {
		return profile, nil, fmt.Errorf("%w: %s", ErrDataLoad, err.Error())
	}

	char := characters[util.Random(0, int64(len(characters)-1))]

	profile.XP = profile.XP + uint(char.Multiplier)
	profile.Active = &char
	// TODO: register in multiverse

	evts := []components.Event{
		components.NewEvent(components.NewCharacterSpawnEvent(char, options...)),
		components.NewEvent(components.NewEarnedXPEvent(uint(char.Multiplier), options...)),
	}

	return profile, evts, nil
}
