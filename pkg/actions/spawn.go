package actions

import (
	"errors"

	"github.com/ciphermountain/deadenz/internal/util"
	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/events"
)

var (
	ErrAlreadySpawnedIn = errors.New("character already exists; cannot spawn in")
)

type WithData struct {
	// base items
	Items      []deadenz.Item
	Characters []deadenz.Character

	// events
	ItemDecisions []events.ItemDecisionEvent
}

// Spawn assigns a new character to an existing profile and modifies xp, backpack,
// and stats. Events emitted include spawn event and earned xp event. Will return
// an already spawned error if profile has an active character.
func (d *WithData) Spawn(profile deadenz.Profile) (deadenz.Profile, []events.Event, error) {
	// short circuit if the user has an active character
	if profile.Active != nil {
		return profile, nil, ErrAlreadySpawnedIn
	}

	char := d.Characters[util.Random(0, int64(len(d.Characters)-1))]
	profile.XP = profile.XP + uint(char.Multiplier)
	profile.Active = &char
	// TODO: register in multiverse

	evts := []events.Event{
		events.NewCharacterSpawnEvent(char),
		events.NewEarnedXPEvent(uint(char.Multiplier)),
	}

	return profile, evts, nil
}
