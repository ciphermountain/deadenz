package components

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
}

type Stats struct {
	Wit   int
	Skill int
	Humor int
}
