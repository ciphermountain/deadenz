package core

// Profile defines the basic properties of an active player
// this is the core struct to pass around
type Profile struct {
	XP        uint
	Currency  uint
	WalkCount uint16
	Active    *Character
}
