package deadenz

import (
	"errors"

	"github.com/ciphermountain/deadenz/internal/util"
	"github.com/ciphermountain/deadenz/pkg/components"
)

var (
	DefaultItemFindRate float64 = 0.5 // % of the time that will result in a findable item
	DefaultDieRate      float64 = 0.3
	ErrNotSpawnedIn             = errors.New("no active character. spawnin to begin") // TODO: breaks multi-language support
	ErrBackpackTooSmall         = errors.New("not enough room in your backpack")      // TODO: breaks multi-language support
)

func Walk(
	profile *components.Profile,
	loader Loader,
	conf Config,
) (*components.Profile, []components.Event, error) {
	if profile.Active == nil {
		return profile, nil, ErrNotSpawnedIn
	}

	which := util.Random(0, 1000)

	var nextFunc func(*components.Profile, Loader, Config) (*components.Profile, []components.Event, error)

	// configured % of the time will result in a findable item
	if which < int64(conf.ItemFindRate*1000) {
		nextFunc = findItem
	} else {
		nextFunc = encounter
	}

	p, evts, err := nextFunc(profile, loader, conf)
	if err != nil {
		return profile, nil, err
	}

	profile = p

	// apply default earnings for all paths
	evts = append(
		evts,
		components.NewEvent(components.NewEarnedXPEvent(uint(profile.Active.Multiplier))),
		components.NewEvent(components.NewEarnedTokenEvent(uint(profile.Active.Multiplier*3))),
	)

	profile.XP = profile.XP + uint(profile.Active.Multiplier)
	profile.Currency = profile.Currency + uint(profile.Active.Multiplier*3)

	return profile, evts, nil
}

func findItem(profile *components.Profile, loader Loader, _ Config) (*components.Profile, []components.Event, error) {
	var items []components.Item
	if err := loader.Load(&items); err != nil {
		return profile, nil, err
	}

	findableItems := make([]components.Item, 0, len(items))
	for idx, item := range items {
		if item.Findable {
			findableItems = append(findableItems, items[idx])
		}
	}

	var decisions []components.ItemDecisionEvent
	if err := loader.Load(&decisions); err != nil {
		return profile, nil, err
	}

	// random item from loaded data
	idx := util.Random(0, int64(len(findableItems)-1))
	randomItem := findableItems[idx]

	evts := []components.Event{
		components.NewEvent(components.NewFindEvent(randomItem)),
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

	return profile, append(evts, components.NewEvent(dec)), nil
}

func encounter(profile *components.Profile, loader Loader, conf Config) (*components.Profile, []components.Event, error) {
	var encounters []components.EncounterEvent
	if err := loader.Load(&encounters); err != nil {
		return profile, nil, err
	}

	evts := []components.Event{
		components.NewEvent(encounters[util.Random(0, int64(len(encounters)-1))]),
	}

	p, e, err := action(profile, loader, conf)
	if err != nil {
		return profile, nil, err
	}

	return p, append(evts, e...), nil
}

func action(profile *components.Profile, loader Loader, conf Config) (*components.Profile, []components.Event, error) {
	var actions []components.ActionEvent
	if err := loader.Load(&actions); err != nil {
		return profile, nil, err
	}

	evts := []components.Event{
		components.NewEvent(actions[util.Random(0, int64(len(actions)-1))]),
	}

	p, e, err := mutation(profile, loader, conf)
	if err != nil {
		return profile, nil, err
	}

	return p, append(evts, e...), nil
}

func mutation(profile *components.Profile, loader Loader, conf Config) (*components.Profile, []components.Event, error) {
	var live []components.LiveMutationEvent
	if err := loader.Load(&live); err != nil {
		return profile, nil, err
	}

	var die []components.DieMutationEvent
	if err := loader.Load(&die); err != nil {
		return profile, nil, err
	}

	evts := []components.Event{
		newRandomMutationEvent(live, die, conf.DeathRate),
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

func newRandomMutationEvent(
	live []components.LiveMutationEvent,
	die []components.DieMutationEvent,
	deathRate float64,
) components.Event {
	if util.Random(0, 1000) < int64(deathRate*1000) {
		return components.NewEvent(die[util.Random(0, int64(len(die)-1))])
	}

	return components.NewEvent(live[util.Random(0, int64(len(live)-1))])
}
