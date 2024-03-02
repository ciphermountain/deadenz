package actions

import (
	"errors"

	"github.com/ciphermountain/deadenz/internal/util"
	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/events"
	"github.com/ciphermountain/deadenz/pkg/items"
)

var (
	ErrNotSpawnedIn     = errors.New("no active character. spawnin to begin")
	ErrBackpackTooSmall = errors.New("not enough room in your backpack")
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
		nextFunc = d.encounter
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
EventLoop:
	for _, evt := range evts {
		switch evt.(type) {
		case *events.DieMutationEvent:
			profile = deathEventMiddleware(profile)

			// TODO: what about conflicting events?
			// short circuit on death
			break EventLoop
		}
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

	var err error

	dec := d.ItemDecisions[util.Random(0, int64(len(d.ItemDecisions)-1))]
	if dec.AddToBackpack() {
		// do the add to backpack
		profile, err = addToBackpackMiddleware(profile, randomItem)
		if err != nil {
			return profile, evts, err
		}
	}

	return profile, append(evts, dec), nil
}

func (d *WithData) encounter(profile deadenz.Profile) (deadenz.Profile, []events.Event, error) {
	evts := []events.Event{
		events.NewRandomEncounterEvent(),
	}

	p, e, err := d.action(profile)
	if err != nil {
		return profile, nil, err
	}

	return p, append(evts, e...), nil
}

func (d *WithData) action(profile deadenz.Profile) (deadenz.Profile, []events.Event, error) {
	evts := []events.Event{
		d.Actions[util.Random(0, int64(len(d.Actions)-1))],
	}

	p, e, err := d.mutation(profile)
	if err != nil {
		return profile, nil, err
	}

	return p, append(evts, e...), nil
}

func (d *WithData) mutation(profile deadenz.Profile) (deadenz.Profile, []events.Event, error) {
	evts := []events.Event{
		events.NewRandomMutationEvent(d.LiveMutations, d.DieMutations, events.DefaultDieRate),
	}

	return profile, evts, nil
}

func deathEventMiddleware(profile deadenz.Profile) deadenz.Profile {
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

	return profile
}

func addToBackpackMiddleware(profile deadenz.Profile, item deadenz.Item) (deadenz.Profile, error) {
	if len(profile.Backpack) < int(profile.BackpackLimit) {
		profile.Backpack = append([]deadenz.Item{item}, profile.Backpack...)
	} else {
		return profile, ErrBackpackTooSmall
	}

	return profile, nil
}
