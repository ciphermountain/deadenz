package middleware

import (
	"errors"
	"fmt"
	"math"
	"time"

	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/opts"
)

var (
	ErrNilProfile  = errors.New("profile cannot be nil")
	ErrWalkTooMuch = errors.New("[ERR-6002] profile has walked too much")
)

type ItemProvider interface {
	Item(components.ItemType, ...opts.Option) (*components.Item, error)
	Items(...opts.Option) ([]components.Item, error)
}

// WalkLimiter applies defined limits to the walk command usage on a profile
func WalkLimiter(hourlyLimit uint16, items ItemProvider) deadenz.PreRunFunc {
	return func(cmd deadenz.CommandType, profile *components.Profile, options ...opts.Option) (*components.Profile, error) {
		if cmd != deadenz.WalkCommandType {
			return profile, nil
		}

		if profile == nil {
			return nil, ErrNilProfile
		}

		limit := int64(hourlyLimit)

		// if active item can extend limit, use it
		if profile.ActiveItem != nil {
			if item, err := items.Item(*profile.ActiveItem, options...); err == nil {
				if item.IsUsable() {
					usable := item.AsUsableItem()

					if usable.ImprovesWalking() {
						limit = int64(1+usable.Efficiency(profile.Stats)) * limit
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

		diff := time.Since(profile.Limits.LastWalk) / time.Millisecond // duration since last walk
		unit := float64(time.Hour/time.Millisecond) / float64(limit)
		jumps := math.Floor(float64(diff) / unit)

		count := profile.Limits.WalkCount - uint64(jumps)

		// check overflow
		if count > profile.Limits.WalkCount {
			count = 0
		}

		profile.Limits.WalkCount = count

		if profile.Limits.WalkCount > uint64(limit) {
			return profile, fmt.Errorf("%w", ErrWalkTooMuch)
		}

		profile.Limits.LastWalk = time.Now()
		profile.Limits.WalkCount++

		return profile, nil
	}
}

func WalkStatBuilder(items ItemProvider, cmds ...deadenz.CommandType) deadenz.PreRunFunc {
	return func(cmd deadenz.CommandType, profile *components.Profile, options ...opts.Option) (*components.Profile, error) {
		var found bool

		for _, c := range cmds {
			if c == cmd {
				found = true
			}
		}

		if !found {
			return profile, nil
		}

		if profile.ActiveItem == nil {
			return profile, nil
		}

		item, err := items.Item(*profile.ActiveItem, options...)
		if err != nil {
			return profile, nil
		}

		// mutation here should only happen walking improves
		if item.IsUsable() && item.Usability.ImprovesWalking {
			profile = item.AsUsableItem().Mutate(profile)
		}

		return profile, nil
	}
}

func WalkDeathEventMiddleware() deadenz.PostRunFunc {
	return func(_ deadenz.CommandType, profile *components.Profile, evts []components.Event, _ ...opts.Option) (*components.Profile, error) {
		if profile.Active == nil {
			return profile, nil
		}

	EventLoop:
		for _, evt := range evts {
			switch evt.Typed().(type) {
			case components.DieMutationEvent, components.TripTrapMutationEvent:
				profile.Active = nil

				break EventLoop
			}
		}

		return profile, nil
	}
}
