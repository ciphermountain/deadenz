package components

type StatMutator struct {
	Stat  string
	Value int
}

func NewStatMutator(stat string, val int) *StatMutator {
	return &StatMutator{
		Stat:  stat,
		Value: val,
	}
}

func (m *StatMutator) Mutate(profile *Profile) *Profile {
	switch m.Stat {
	case "wit":
		profile.Stats.Wit += m.Value
	case "skill":
		profile.Stats.Skill += m.Value
	case "humor":
		profile.Stats.Humor += m.Value
	}

	return profile
}

type BackpackLimitMutator struct {
	Limit uint8
}

func NewBackpackLimitMutator(limit uint8) *BackpackLimitMutator {
	return &BackpackLimitMutator{
		Limit: limit,
	}
}

func (m *BackpackLimitMutator) Mutate(profile *Profile) *Profile {
	if len(profile.Backpack) > int(m.Limit) {
		profile.Backpack = profile.Backpack[:m.Limit]
	}

	profile.BackpackLimit = uint8(m.Limit)

	return profile
}
