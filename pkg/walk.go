package deadenz

import (
	"errors"

	"github.com/ciphermountain/deadenz/internal/util"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/events"
)

var (
	ErrNotSpawnedIn     = errors.New("no active character. spawnin to begin")
	ErrBackpackTooSmall = errors.New("not enough room in your backpack")
)

func Walk(profile *components.Profile, loader Loader) (*components.Profile, []components.Event, error) {
	if profile.Active == nil {
		return profile, nil, ErrNotSpawnedIn
	}

	which := util.Random(0, 100)

	var nextFunc func(*components.Profile, Loader) (*components.Profile, []components.Event, error)

	// 35% of the time will result in a findable item
	if which < 35 {
		nextFunc = findItem
	} else {
		nextFunc = encounter
	}

	p, evts, err := nextFunc(profile, loader)
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

	return profile, evts, nil
}

func findItem(profile *components.Profile, loader Loader) (*components.Profile, []components.Event, error) {
	var items []components.Item
	if err := loader.Load(&items); err != nil {
		return profile, nil, err
	}

	var decisions []events.ItemDecisionEvent
	if err := loader.Load(&decisions); err != nil {
		return profile, nil, err
	}

	// random item from loaded data
	idx := util.Random(0, int64(len(items)-1))
	randomItem := items[idx]

	evts := []components.Event{
		events.NewFindEvent(randomItem),
	}

	var err error

	dec := decisions[util.Random(0, int64(len(decisions)-1))]
	if dec.AddToBackpack() {
		// do the add to backpack
		profile, err = addToBackpack(profile, randomItem)
		if err != nil {
			// the only possible error here is the backpack being too small
			// the event needs to be surfaced, but the profile should be allowed to be modified
			return profile, evts, err
		}
	}

	return profile, append(evts, dec), nil
}

func encounter(profile *components.Profile, loader Loader) (*components.Profile, []components.Event, error) {
	var encounters []events.EncounterEvent
	if err := loader.Load(&encounters); err != nil {
		return profile, nil, err
	}

	evts := []components.Event{
		encounters[util.Random(0, int64(len(encounters)-1))],
	}

	p, e, err := action(profile, loader)
	if err != nil {
		return profile, nil, err
	}

	return p, append(evts, e...), nil
}

func action(profile *components.Profile, loader Loader) (*components.Profile, []components.Event, error) {
	var actions []events.ActionEvent
	if err := loader.Load(&actions); err != nil {
		return profile, nil, err
	}

	evts := []components.Event{
		actions[util.Random(0, int64(len(actions)-1))],
	}

	p, e, err := mutation(profile, loader)
	if err != nil {
		return profile, nil, err
	}

	return p, append(evts, e...), nil
}

func mutation(profile *components.Profile, loader Loader) (*components.Profile, []components.Event, error) {
	var live []events.LiveMutationEvent
	if err := loader.Load(&live); err != nil {
		return profile, nil, err
	}

	var die []events.DieMutationEvent
	if err := loader.Load(&die); err != nil {
		return profile, nil, err
	}

	evts := []components.Event{
		events.NewRandomMutationEvent(live, die, events.DefaultDieRate),
	}

	return profile, evts, nil
}

func addToBackpack(profile *components.Profile, item components.Item) (*components.Profile, error) {
	if len(profile.Backpack) < int(profile.BackpackLimit) {
		profile.Backpack = append([]components.ItemType{item.Type}, profile.Backpack...)
	} else {
		return profile, ErrBackpackTooSmall
	}

	return profile, nil
}
