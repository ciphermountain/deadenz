package deadenz

import (
	"errors"

	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/events"
)

var ErrUnrecognizedCommand = errors.New("unrecognized command")

// Result represents the state change of applying one step of the game on a player profile.
type Result struct {
	DefaultCmd CommandType
	Profile    *components.Profile
	Events     []components.Event
}

// PreRunFunc can read a profile, modify and return it.
type PreRunFunc func(CommandType, *components.Profile) (*components.Profile, error)

// PreRunFunc can read a profile with events, modify the profile, and return it.
type PostRunFunc func(CommandType, *components.Profile, []components.Event) (*components.Profile, error)

func RunActionCommand(
	command CommandType,
	profile *components.Profile,
	loader Loader,
	preRun []PreRunFunc,
	postRun []PostRunFunc,
) (Result, error) {
	if profile == nil {
		return Result{}, errors.New("profile required")
	}

	original := *profile
	step := Result{
		Profile: profile,
	}

	for idx := range preRun {
		var err error

		step.Profile, err = preRun[idx](command, step.Profile)
		if err != nil {
			return Result{Profile: &original}, err
		}
	}
	// return profile, errors.New("you have walked too much. you need to rest for an hour. there's a park bench nearby")

	switch command {
	case SpawninCommandType:
		var err error

		step.Profile, step.Events, err = Spawn(step.Profile, loader)
		if err != nil {
			return Result{Profile: &original}, err
		}

		step.DefaultCmd = WalkCommandType
	case WalkCommandType:
		var err error

		step.Profile, step.Events, err = Walk(step.Profile, loader)
		if err != nil {
			if !errors.Is(err, ErrBackpackTooSmall) {
				return Result{Profile: &original}, err
			}

			// TODO: make a better message
			step.Events = append(step.Events, events.NewItemDecisionEvent("your backpack is too small"))
			err = nil
		}

		step.DefaultCmd = WalkCommandType

		if profile.Active == nil {
			step.DefaultCmd = SpawninCommandType
		}
	default:
		return step, ErrUnrecognizedCommand
	}

	for idx := range postRun {
		var err error

		step.Profile, err = postRun[idx](command, step.Profile, step.Events)
		if err != nil {
			return Result{Profile: &original}, err
		}
	}

	return step, nil
}
