package parse

import (
	"encoding/json"

	"github.com/ciphermountain/deadenz/pkg/components"
)

func ItemsFromJSON(b []byte) ([]components.Item, error) {
	type basicItem struct {
		Name string
	}

	var loaded []basicItem

	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, err
	}

	items := []components.Item{
		components.NewLocker(),
		components.NewWalkingStick(),
	}

	for i, l := range loaded {
		items = append(items, components.Item{
			Type: components.ItemType(i + 3),
			Name: l.Name,
			Mutate: func(s components.Stats) components.Stats {
				return s
			},
		})
	}

	return items, nil
}
