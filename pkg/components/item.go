package components

type ItemType uint64

type Item struct {
	Type   ItemType    `json:"type"`
	Name   string      `json:"name"`
	Mutate MutatorFunc `json:"-"`
}

type MutatorFunc func(Stats) Stats

const (
	Locker ItemType = iota + 1
	WalkingStick
)

func NewLocker() Item {
	return Item{
		Type: Locker,
		Name: "a locker",
		Mutate: func(s Stats) Stats {
			return s
		},
	}
}

func NewWalkingStick() Item {
	return Item{
		Type: WalkingStick,
		Name: "a walking stick",
		Mutate: func(s Stats) Stats {
			return s
		},
	}
}
