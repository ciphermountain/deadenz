package deadenz

import (
	"errors"

	"github.com/ciphermountain/deadenz/pkg/components"
)

var ErrUnrecognizedCommand = errors.New("unrecognized command")

// Result represents the state change of applying one step of the game on a player profile.
type Result struct {
	DefaultCmd CommandType
	Profile    components.Profile
	Events     []components.Event
}

// PreRunFunc can read a profile, modify and return it.
type PreRunFunc func(CommandType, components.Profile) (components.Profile, error)

// PreRunFunc can read a profile with events, modify the profile, and return it.
type PostRunFunc func(CommandType, components.Profile, []components.Event) (components.Profile, error)

func RunActionCommand(
	command CommandType,
	profile components.Profile,
	loader Loader,
	preRun []PreRunFunc,
	postRun []PostRunFunc,
) (Result, error) {
	var step Result

	for idx := range preRun {
		var err error

		profile, err = preRun[idx](command, profile)
		if err != nil {
			return step, err
		}
	}

	switch command {
	case SpawninCommandType:
		var err error

		step.Profile, step.Events, err = Spawn(profile, loader)
		if err != nil {
			return step, err
		}

		step.DefaultCmd = WalkCommandType
	case WalkCommandType:
		var err error

		step.Profile, step.Events, err = Walk(profile, loader)
		if err != nil {
			return step, err
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
			return step, err
		}
	}

	return step, nil
}
