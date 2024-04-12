package deadenz

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/events"
	proto "github.com/ciphermountain/deadenz/pkg/proto/multiverse"
	"github.com/ciphermountain/deadenz/pkg/service/multiverse"
)

var ErrUnrecognizedCommand = errors.New("unrecognized command")

// Result represents the state change of applying one step of the game on a player profile.
type Result struct {
	DefaultCmd CommandType
	Profile    components.Profile
	Events     []components.Event
}

func RunActionCommand(command CommandType, profile components.Profile, loader Loader, client *multiverse.Client) (Result, error) {
	var step Result

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

		if client != nil {
			publishEvents(step.Events, client)
		}

		step.DefaultCmd = WalkCommandType

		if profile.Active == nil {
			step.DefaultCmd = SpawninCommandType
		}
	default:
		return step, ErrUnrecognizedCommand
	}

	return step, nil
}

func publishEvents(evts []components.Event, client *multiverse.Client) {
	for _, evt := range evts {
		switch evt.(type) {
		case events.DieMutationEvent, events.CharacterSpawnEvent: // only spawn and die events are supported
			bts, err := json.Marshal(evt)
			if err != nil {
				continue
			}

			client.PublishEvent(context.Background(), &proto.Event{
				Data: bts,
			})
		}
	}
}
