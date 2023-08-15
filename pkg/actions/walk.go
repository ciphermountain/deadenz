package actions

import (
	"errors"

	"github.com/ciphermountain/deadenz/internal/util"
	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/events"
	"github.com/ciphermountain/deadenz/pkg/items"
)

var (
	ErrNotSpawnedIn     = errors.New("no active character. spawnin to begin.")
	ErrBackpackTooSmall = errors.New("not enough room in your backpack.")
)

func (d *WithData) Walk(profile deadenz.Profile) (deadenz.Profile, []events.Event, error) {
	if profile.Active == nil {
		return profile, nil, ErrNotSpawnedIn
	}

	if profile.ActiveItem != nil && profile.ActiveItem.Type == items.WalkingStick {
		profile.Stats = profile.ActiveItem.Mutate(profile.Stats)
	}

	// TODO: apply multiverse death filter

	which := util.Random(0, 100)

	var nextFunc func(profile deadenz.Profile) (deadenz.Profile, []events.Event, error)

	// 35% of the time will result in a findable item
	if which < 35 {
		nextFunc = d.findItem
	} else {
		nextFunc = encounter
	}

	p, evts, err := nextFunc(profile)
	if err != nil {
		return profile, nil, err
	}

	profile = p

	// apply default earnings for all paths
	evts = append(
		evts,
		events.NewEarnedXPEvent(uint(profile.Active.Multiplier)),
		events.NewEarnedTokenEvent(uint(profile.Active.Multiplier*3)),
	)

	profile.XP = profile.XP + uint(profile.Active.Multiplier)
	profile.Currency = profile.Currency + uint(profile.Active.Multiplier*3)

	// if any event is a death event, remove active character and apply backpack recovery
	for _, evt := range evts {
		if _, ok := evt.(*events.DieMutationEvent); ok {
			profile.Active = nil

			backpack := []deadenz.Item{}

			// a locker allows the backpack items to be recovered
			if profile.ActiveItem != nil && profile.ActiveItem.Type == items.Locker {
				profile.Stats = profile.ActiveItem.Mutate(profile.Stats)
				profile.ActiveItem = nil

				// backpack recovery
				backpack = profile.Backpack
			}

			profile.Backpack = backpack

			break
		}
	}

	return profile, evts, nil
}

func encounter(profile deadenz.Profile) (deadenz.Profile, []events.Event, error) {
	evts := []events.Event{
		events.NewRandomEncounterEvent(),
	}

	p, e, err := action(profile)
	if err != nil {
		return profile, nil, err
	}

	return p, append(evts, e...), nil
}

func action(profile deadenz.Profile) (deadenz.Profile, []events.Event, error) {
	evts := []events.Event{
		events.NewRandomActionEvent(),
	}

	p, e, err := mutation(profile)
	if err != nil {
		return profile, nil, err
	}

	return p, append(evts, e...), nil
}

func mutation(profile deadenz.Profile) (deadenz.Profile, []events.Event, error) {
	evts := []events.Event{
		events.NewRandomMutationEvent(),
	}

	return profile, evts, nil
}

func (d *WithData) findItem(profile deadenz.Profile) (deadenz.Profile, []events.Event, error) {
	// random item from loaded data
	idx := util.Random(0, int64(len(d.Items)-1))
	randomItem := d.Items[idx]

	evts := []events.Event{
		events.NewFindEvent(randomItem),
	}

	p, e, err := decision(profile)
	if err != nil {
		return profile, nil, err
	}

	// TODO: don't do string comparisons on events
	if len(e) > 0 && e[0].String() == "add it to your backpack" {
		if len(profile.Backpack) < int(profile.BackpackLimit) {
			profile.Backpack = append([]deadenz.Item{randomItem}, profile.Backpack...)
		} else {
			return p, evts, ErrBackpackTooSmall
		}
	}

	return p, append(evts, e...), nil
}

func decision(profile deadenz.Profile) (deadenz.Profile, []events.Event, error) {
	dec := events.NewRandomDecisionEvent()

	return profile, []events.Event{dec}, nil
}
