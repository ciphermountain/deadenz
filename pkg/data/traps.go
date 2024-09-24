package data

import (
	"errors"

	"github.com/ciphermountain/deadenz/pkg/components"
)

type TrapProvider struct {
	loader *DataLoader
}

func NewTrapProviderFromLoader(loader *DataLoader) *TrapProvider {
	return &TrapProvider{
		loader: loader,
	}
}

func (p *TrapProvider) Traps(profile *components.Profile) ([]components.Trap, error) {
	var traps []components.Trap
	if err := p.loader.Load(&traps); err != nil {
		return nil, err
	}

	// filter out traps set by the provided profile
	set := make([]components.Trap, 0, len(traps))
	for _, trap := range traps {
		if trap.SetBy != profile.UUID {
			set = append(set, trap)
		}
	}

	return set, nil
}

func (p *TrapProvider) TripRandom(*components.Profile) (components.Trap, error) {
	return components.Trap{}, errors.New("unimplemented")
}
