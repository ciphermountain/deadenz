package middleware

import (
	"context"
	"encoding/json"

	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/opts"
)

type EventPublisher interface {
	PublishGameEvent(context.Context, string, []byte) error
}

func PublishEventsToMultiverse(client EventPublisher) deadenz.PostRunFunc {
	return func(cmd deadenz.CommandType, profile *components.Profile, evts []components.Event, _ ...opts.Option) (*components.Profile, error) {
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

func publishEvents(profile *components.Profile, evts []components.Event, client EventPublisher) {
	for _, evt := range evts {
		switch typed := evt.Typed().(type) {
		case components.DieMutationEvent:
			_ = marshalAndSend(components.NewDieMutationEventWithCharacter(*profile.Active, typed), client, profile.UUID)
		case components.CharacterSpawnEvent: // only spawn and die events are supported
			_ = marshalAndSend(typed, client, profile.UUID)
		default:
			continue
		}
	}
}

func marshalAndSend[T any](evt T, client EventPublisher, id string) error {
	bts, err := json.Marshal(evt)
	if err != nil {
		return err
	}

	return client.PublishGameEvent(context.Background(), id, bts)
}
