package deadenz

import (
	"errors"

	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/opts"
)

var ErrUnrecognizedCommand = errors.New("unrecognized command")

// Result represents the state change of applying one step of the game on a player profile.
type Result struct {
	DefaultCmd CommandType
	Profile    *components.Profile
	Events     []components.Event
}

// PreRunFunc can read a profile, modify and return it.
type PreRunFunc func(CommandType, *components.Profile, ...opts.Option) (*components.Profile, error)

// PreRunFunc can read a profile with events, modify the profile, and return it.
type PostRunFunc func(CommandType, *components.Profile, []components.Event, ...opts.Option) (*components.Profile, error)

func RunActionCommand(
	command CommandType,
	profile *components.Profile,
	loader Loader,
	conf Config,
	preRun []PreRunFunc,
	postRun []PostRunFunc,
	options ...opts.Option,
) (Result, error) {
	if profile == nil {
		return Result{}, errors.New("profile required")
	}

	original := *profile
	step := Result{
		Profile: profile,
	}

PreRun:
	for idx := range preRun {
		var err error

		step.Profile, err = preRun[idx](command, step.Profile, options...)
		if err != nil {
			// if a trap is tripped, switch the command type
			var trapErr ErrTrapTripped
			if errors.As(err, &trapErr) {
				// create a special death event and set a bypass command type
				step.Events = append(step.Events, components.NewTripTrapMutationEvent(trapErr.Trap.Message))
				command = TrapCommandType

				break PreRun
			}

			return Result{Profile: &original}, err
		}
	}

	switch command {
	case SpawninCommandType:
		var err error

		step.Profile, step.Events, err = Spawn(step.Profile, loader, conf, options...)
		if err != nil {
			return Result{Profile: &original}, err
		}

		step.DefaultCmd = WalkCommandType
	case WalkCommandType:
		var err error

		step.Profile, step.Events, err = Walk(step.Profile, loader, conf, options...)
		if err != nil {
			if !errors.Is(err, ErrBackpackTooSmall) {
				return Result{Profile: &original}, err
			}

			// TODO: make a better message
			step.Events = append(step.Events, components.NewEvent(components.NewItemDecisionEvent("your backpack is too small")))
			err = nil
		}

		step.DefaultCmd = WalkCommandType
	case TrapCommandType:
		step.DefaultCmd = SpawninCommandType
	default:
		return step, ErrUnrecognizedCommand
	}

	for idx := range postRun {
		var err error

		step.Profile, err = postRun[idx](command, step.Profile, step.Events, options...)
		if err != nil {
			return Result{Profile: &original}, err
		}
	}

	if profile.Active == nil {
		step.DefaultCmd = SpawninCommandType
	}

	return step, nil
}
