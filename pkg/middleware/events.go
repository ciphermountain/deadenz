package middleware

import (
	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/events"
)

// DeathActiveItemMiddleware applies the mutation of an active item as long as a death event exists and the
// active item matches the provided item type. The active item is removed after the mutation is applied.
func DeathActiveItemMiddleware(items ItemProvider) deadenz.PostRunFunc {
	return func(_ deadenz.CommandType, profile *components.Profile, evts []components.Event) (*components.Profile, error) {
		// if any event is a death event, remove active character and apply backpack recovery
	EventLoop:
		for _, evt := range evts {
			switch evt.(type) {
			case events.DieMutationEvent:
				// always remove active item
				temp := profile.ActiveItem
				profile.ActiveItem = nil

				if temp != nil {
					item, err := items.Item(*temp)
					if err != nil {
						return profile, nil
					}

					if item.IsUsable() {
						profile = item.AsUsableItem().ModifyBackpackContents(item.Mutate(profile))
					} else {
						profile.Backpack = []components.ItemType{}
					}
				} else {
					profile.Backpack = []components.ItemType{}
				}

				break EventLoop
			}
		}

		return profile, nil
	}
}
