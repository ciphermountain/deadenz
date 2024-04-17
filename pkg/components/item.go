package components

type ItemType uint64

type MutatorFunc func(Profile) Profile

type Item struct {
	Type      ItemType
	Name      string
	Usability *Usability
	Mutators  []MutatorFunc
}

type Usability struct {
	ImprovesWalking   bool  `json:"improves_walking"`
	SaveBackpackItems uint8 `json:"save_backpack_items"`
}

func (i Item) Mutate(profile Profile) Profile {
	for _, f := range i.Mutators {
		profile = f(profile)
	}

	return profile
}

func (i Item) IsUsable() bool {
	return i.Usability != nil
}

func (i Item) AsUsableItem() UsableItem {
	return newUsableItem(i)
}

type UsableItem struct {
	item Item
}

func newUsableItem(item Item) UsableItem {
	return UsableItem{item: item}
}

func (i UsableItem) ImprovesWalking() bool {
	return i.item.Usability.ImprovesWalking
}

func (i UsableItem) ModifyBackpackContents(profile Profile) Profile {
	limit := i.item.Usability.SaveBackpackItems

	if limit > profile.BackpackLimit {
		limit = profile.BackpackLimit
	}

	if int(limit) < len(profile.Backpack) {
		profile.Backpack = profile.Backpack[:limit]
	}

	return profile
}

func (i UsableItem) Efficiency(stats Stats) int {
	/*
		const x = stats.skill;
		const scale = 10000;

		return (x*x)/((x*x)+scale)
	*/

	return 1
}

const (
	Locker ItemType = iota + 1
)

func NewLocker() Item {
	return Item{
		Type:     Locker,
		Name:     "a locker",
		Mutators: []MutatorFunc{MutateSkillBy(1)},
	}
}

func MutateWitBy(val int) MutatorFunc {
	return func(profile Profile) Profile {
		profile.Stats.Wit += val

		return profile
	}
}

func MutateSkillBy(val int) MutatorFunc {
	return func(profile Profile) Profile {
		profile.Stats.Skill += val

		return profile
	}
}

func MutateHumorBy(val int) MutatorFunc {
	return func(profile Profile) Profile {
		profile.Stats.Humor += val

		return profile
	}
}

func BackpackLimitMutator(limit uint8) MutatorFunc {
	return func(profile Profile) Profile {
		if len(profile.Backpack) > int(limit) {
			profile.Backpack = profile.Backpack[:limit]
		}

		profile.BackpackLimit = uint8(limit)

		return profile
	}
}
