package middleware

import (
	"context"
	"encoding/json"

	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/events"
	service "github.com/ciphermountain/deadenz/pkg/service/multiverse"
)

func PublishEventsToMultiverse(client *service.Client) deadenz.PostRunFunc {
	return func(cmd deadenz.CommandType, profile *components.Profile, evts []components.Event) (*components.Profile, error) {
		// passthrough if not walk or spawnin command
		if cmd != deadenz.WalkCommandType && cmd != deadenz.SpawninCommandType {
			return profile, nil
		}

		if client != nil {
			publishEvents(profile, evts, client)
		}

		return profile, nil
	}
}

func publishEvents(profile *components.Profile, evts []components.Event, client *service.Client) {
	for _, evt := range evts {
		switch typed := evt.(type) {
		case events.DieMutationEvent:
			_ = marshalAndSend(events.NewDieMutationEventWithCharacter(*profile.Active, typed), client, profile.UUID)
		case events.CharacterSpawnEvent: // only spawn and die events are supported
			_ = marshalAndSend(typed, client, profile.UUID)
		default:
			continue
		}
	}
}

func marshalAndSend[T any](evt T, client *service.Client, id string) error {
	bts, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	return client.PublishGameEvent(context.Background(), id, bts)
}
