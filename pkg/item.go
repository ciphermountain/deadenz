package deadenz

type ItemType uint64

type Item struct {
	Type   ItemType    `json:"type"`
	Name   string      `json:"name"`
	Mutate MutatorFunc `json:"-"`
}

type MutatorFunc func(Stats) Stats
