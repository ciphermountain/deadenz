package core

import (
	"fmt"
	"errors"
)

var (
	ErrAlreadySpawnedIn = errors.New("character already exists; cannot spawn in")
	ErrMissingProfile = errors.New("profile required to run this function")
)

func SpawnIn(p *Profile) (*Profile, error) {
	if p == nil {
		return nil, fmt.Errorf("%w: spawnin", ErrMissingProfile)
	}

	return p, nil
}
