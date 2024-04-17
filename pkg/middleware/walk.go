package middleware

import (
	"errors"
	"time"

	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/components"
)

const (
	hourInMilliseconds uint64 = 1000 * 60 * 60
)

type ItemProvider interface {
	Item(components.ItemType) (*components.Item, error)
	Items() ([]components.Item, error)
}

// WalkLimiter applies defined limits to the walk command usage on a profile
func WalkLimiter(hourlyLimit uint64, items ItemProvider) deadenz.PreRunFunc {
	return func(cmd deadenz.CommandType, profile components.Profile) (components.Profile, error) {
		if cmd != deadenz.WalkCommandType {
			return profile, nil
		}

		limit := hourlyLimit

		// if active item can extend limit, use it
		if profile.ActiveItem != nil {
			if item, err := items.Item(*profile.ActiveItem); err == nil {
				if item.IsUsable() {
					usable := item.AsUsableItem()

					if usable.ImprovesWalking() {
						limit = uint64((1 + usable.Efficiency(profile.Stats))) * limit
					}
				}
			}
		}

		if profile.Limits == nil {
			profile.Limits = &components.Limits{
				LastWalk:  time.Now(),
				WalkCount: 1,
			}

			return profile, nil
		}

		diff := time.Since(profile.Limits.LastWalk) // duration since last walk
		unit := hourInMilliseconds / limit
		jumps := uint64(diff) / unit

		count := profile.Limits.WalkCount - jumps

		// check overflow
		if count > profile.Limits.WalkCount {
			count = 0
		}

		profile.Limits.WalkCount = count

		if profile.Limits.WalkCount > limit {
			return profile, errors.New("you have walked too much. you need to rest for an hour. there's a park bench nearby")
		}

		return profile, nil
	}
}

func WalkStatBuilder(it components.ItemType, items ItemProvider, cmds ...deadenz.CommandType) deadenz.PreRunFunc {
	return func(cmd deadenz.CommandType, profile components.Profile) (components.Profile, error) {
		var found bool

		for _, c := range cmds {
			if c == cmd {
				found = true
			}
		}

		if !found {
			return profile, nil
		}

		if profile.ActiveItem == nil || *profile.ActiveItem != it {
			return profile, nil
		}

		if item, err := items.Item(it); err == nil {
			profile = item.Mutate(profile)
		}

		return profile, nil
	}
}
