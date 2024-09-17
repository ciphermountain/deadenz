package parse

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/ciphermountain/deadenz/pkg/components"
)

type simpleJSONItem struct {
	Name        string                `json:"name"`
	Findable    bool                  `json:"findable"`
	Purchasable bool                  `json:"purchasable"`
	Price       int64                 `json:"price"`
	Usability   *components.Usability `json:"usability,omitempty"`
}

type decodingJSONItem struct {
	simpleJSONItem
	Mutators []json.RawMessage `json:"mutators,omitempty"`
}

type typer struct {
	Type string `json:"type"`
}

func ItemsFromJSON(b []byte) ([]components.Item, error) {
	var loaded []decodingJSONItem
	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, err
	}

	items := make([]components.Item, len(loaded))

	for idx, item := range loaded {
		mutators := make([]components.ProfileMutator, len(item.Mutators))
		for idx, conf := range item.Mutators {
			var typed typer
			if err := json.Unmarshal(conf, &typed); err != nil {
				return nil, err
			}

			var (
				mutator func([]byte) (components.ProfileMutator, error)
				err     error
			)

			switch typed.Type {
			case "stats":
				mutator = asStatMutator
			case "backpack_limit":
				mutator = asBackpackLimitMutator
			default:
				return nil, errors.New("unrecognized mutator type")
			}

			if mutators[idx], err = mutator(conf); err != nil {
				return nil, err
			}
		}

		items[idx] = components.Item{
			Type:        components.ItemType(idx + 1),
			Name:        item.Name,
			Findable:    item.Findable,
			Purchasable: item.Purchasable,
			Price:       item.Price,
			Usability:   item.Usability,
			Mutators:    mutators,
		}
	}

	return items, nil
}

func asStatMutator(data []byte) (components.ProfileMutator, error) {
	type jsonStatMutator struct {
		StatName string `json:"stat_name"`
		Mutation string `json:"mutation"`
	}

	var statMut jsonStatMutator
	if err := json.Unmarshal(data, &statMut); err != nil {
		return nil, err
	}

	value, err := strconv.Atoi(statMut.Mutation)
	if err != nil {
		return nil, err
	}

	switch statMut.StatName {
	case "wit":
		return components.NewStatMutator("wit", value), nil
	case "skill":
		return components.NewStatMutator("skill", value), nil
	case "humor":
		return components.NewStatMutator("humor", value), nil
	default:
		return nil, errors.New("invalid stat name")
	}
}

func asBackpackLimitMutator(data []byte) (components.ProfileMutator, error) {
	type mutator struct {
		Limit uint8 `json:"limit"`
	}

	var mut mutator
	if err := json.Unmarshal(data, &mut); err != nil {
		return nil, err
	}

	return components.NewBackpackLimitMutator(mut.Limit), nil
}
