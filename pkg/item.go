package deadenz

type ItemType uint64

type Item struct {
	Type   ItemType
	Name   string
	Mutate MutatorFunc
}

type MutatorFunc func(Stats) Stats
