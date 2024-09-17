package components

import (
	"math"
)

type ItemType uint64

type ProfileMutator interface {
	Mutate(*Profile) *Profile
}

type Item struct {
	Type        ItemType
	Name        string
	Findable    bool
	Purchasable bool
	Price       int64
	Usability   *Usability
	Mutators    []ProfileMutator
}

type Usability struct {
	ImprovesWalking   bool       `json:"improves_walking,omitempty"`
	SaveBackpackItems uint8      `json:"save_backpack_items,omitempty"`
	Efficiency        Efficiency `json:"efficiency,omitempty"`
}

type Efficiency struct {
	Stat  string `json:"stat_name"`
	Scale uint32 `json:"scale"`
}

func (i Item) Mutate(profile *Profile) *Profile {
	for _, f := range i.Mutators {
		profile = f.Mutate(profile)
	}

	return profile
}

func (i Item) IsUsable() bool {
	return i.Usability != nil
}

func (i Item) AsUsableItem() UsableItem {
	return NewUsableItem(i)
}

type efficiencyFunc func(Stats) int

type UsableItem struct {
	item           Item
	efficiencyFunc efficiencyFunc
}

func NewUsableItem(item Item) UsableItem {
	var effFunc efficiencyFunc

	switch item.Usability.Efficiency.Stat {
	case "wit":
		effFunc = ScaledEfficiency(item.Usability.Efficiency.Scale, forWit)
	case "skill":
		effFunc = ScaledEfficiency(item.Usability.Efficiency.Scale, forSkill)
	case "humor":
		effFunc = ScaledEfficiency(item.Usability.Efficiency.Scale, forHumor)
	default:
		effFunc = DefaultEfficiency
	}

	return UsableItem{item: item, efficiencyFunc: effFunc}
}

func (i UsableItem) ImprovesWalking() bool {
	return i.item.Usability.ImprovesWalking
}

func (i UsableItem) ModifyBackpackContents(profile *Profile) *Profile {
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
	return i.efficiencyFunc(stats)
}

func DefaultEfficiency(_ Stats) int {
	return 1
}

func ScaledEfficiency(scale uint32, forStat func(Stats) int) efficiencyFunc {
	return func(stats Stats) int {
		skill := float64(forStat(stats))
		result := math.Ceil((skill * skill) / ((skill * skill) + float64(scale)))

		return int(result)
	}
}

func forWit(stats Stats) int {
	return stats.Wit
}

func forSkill(stats Stats) int {
	return stats.Skill
}

func forHumor(stats Stats) int {
	return stats.Humor
}
