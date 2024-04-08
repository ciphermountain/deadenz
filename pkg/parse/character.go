package parse

import (
	"encoding/json"

	"github.com/ciphermountain/deadenz/pkg/components"
)

func CharactersFromJSON(b []byte) ([]components.Character, error) {
	type basicCharacter struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Mult int    `json:"multiplier"`
	}

	var loaded []basicCharacter

	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, err
	}

	chars := []components.Character{}

	for _, l := range loaded {
		chars = append(chars, components.Character{
			Type:       components.CharacterType(l.Type),
			Name:       l.Name,
			Multiplier: uint8(l.Mult),
		})
	}

	return chars, nil
}
