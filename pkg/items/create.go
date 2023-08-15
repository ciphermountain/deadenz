package items

import (
	"encoding/json"

	deadenz "github.com/ciphermountain/deadenz/pkg"
)

const (
	Locker       deadenz.ItemType = 1
	WalkingStick                  = 2
)

func NewLocker() deadenz.Item {
	return deadenz.Item{
		Type: Locker,
		Name: "a locker",
		Mutate: func(s deadenz.Stats) deadenz.Stats {
			return s
		},
	}
}

func NewWalkingStick() deadenz.Item {
	return deadenz.Item{
		Type: WalkingStick,
		Name: "a walking stick",
		Mutate: func(s deadenz.Stats) deadenz.Stats {
			return s
		},
	}
}

func LoadItems(b []byte) ([]deadenz.Item, error) {
	type basicItem struct {
		Name string
	}

	var loaded []basicItem

	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, err
	}

	items := []deadenz.Item{
		NewLocker(),
		NewWalkingStick(),
	}

	for i, l := range loaded {
		items = append(items, deadenz.Item{
			Type: deadenz.ItemType(i + 3),
			Name: l.Name,
			Mutate: func(s deadenz.Stats) deadenz.Stats {
				return s
			},
		})
	}

	return items, nil
}
