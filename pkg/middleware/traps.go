package middleware

import (
	"errors"

	"github.com/ciphermountain/deadenz/internal/util"
	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/components"
)

type TrapProvider interface {
	// TripRandom should provide an atomic operation that selects and saves that trap was tripped.
	// The user Profile is provided to ensure that a user doesn't trip their own traps.
	TripRandom(*components.Profile) (components.Trap, error)
}

var ErrNoAvailableTraps = errors.New("no available traps")

func PreRunTrapTripper(traps TrapProvider, tripRate float64) deadenz.PreRunFunc {
	return func(cmd deadenz.CommandType, profile *components.Profile) (*components.Profile, error) {
		if cmd != deadenz.WalkCommandType {
			return profile, nil
		}

		which := util.Random(0, 1000)
		if which < int64(tripRate*1000) {
			trap, err := traps.TripRandom(profile)
			if err != nil {
				if errors.Is(err, ErrNoAvailableTraps) {
					return profile, nil
				}

				return nil, err
			}

			// returning an error here stops processing and allows a death event
			// that doesn't fall into post run functions
			return profile, deadenz.ErrTrapTripped{
				Trap: trap,
			}
		}

		return profile, nil
	}
}
