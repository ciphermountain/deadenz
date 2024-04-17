package components

import "time"

// Profile defines the basic properties of an active player
// this is the core struct to pass around
type Profile struct {
	UUID          string
	XP            uint
	Currency      uint
	Active        *Character
	ActiveItem    *ItemType
	BackpackLimit uint8
	Backpack      []ItemType
	Stats         Stats
	Limits        *Limits
}

type Stats struct {
	Wit   int
	Skill int
	Humor int
}

type Limits struct {
	LastWalk  time.Time
	WalkCount uint64
}
