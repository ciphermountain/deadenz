package middleware

import (
	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/events"
)

// DeathActiveItemMiddleware applies the mutation of an active item as long as a death event exists and the
// active item matches the provided item type. The active item is removed after the mutation is applied.
func DeathActiveItemMiddleware(it components.ItemType, items ItemProvider) deadenz.PostRunFunc {
	return func(_ deadenz.CommandType, profile *components.Profile, evts []components.Event) (*components.Profile, error) {
		if profile.ActiveItem == nil || *profile.ActiveItem != it {
			return profile, nil
		}

		// if any event is a death event, remove active character and apply backpack recovery
	EventLoop:
		for _, evt := range evts {
			switch evt.(type) {
			case events.DieMutationEvent:
				item, err := items.Item(*profile.ActiveItem)
				if err != nil {
					return profile, nil
				}

				profile = applyItemDeathEvent(profile, item)

				break EventLoop
			}
		}

		return profile, nil
	}
}

func applyItemDeathEvent(profile *components.Profile, item *components.Item) *components.Profile {
	profile = item.Mutate(profile)

	if item.IsUsable() {
		profile = item.AsUsableItem().ModifyBackpackContents(profile)
	}

	profile.ActiveItem = nil

	return profile
}
