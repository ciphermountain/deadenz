package middleware

import (
	"context"
	"encoding/json"

	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/events"
	proto "github.com/ciphermountain/deadenz/pkg/proto/multiverse"
	service "github.com/ciphermountain/deadenz/pkg/service/multiverse"
)

func PublishEventsToMultiverse(client *service.Client) deadenz.PostRunFunc {
	return func(cmd deadenz.CommandType, profile components.Profile, evts []components.Event) (components.Profile, error) {
		// passthrough if not walk or spawnin command
		if cmd != deadenz.WalkCommandType && cmd != deadenz.SpawninCommandType {
			return profile, nil
		}

		if client != nil {
			publishEvents(evts, client)
		}

		return profile, nil
	}
}

func MultiverseDeathFilter() deadenz.PreRunFunc {
	return func(ct deadenz.CommandType, p components.Profile) (components.Profile, error) {
		return p, nil
	}
}

func publishEvents(evts []components.Event, client *service.Client) {
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
