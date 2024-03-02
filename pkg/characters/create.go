package characters

import (
	"encoding/json"

	deadenz "github.com/ciphermountain/deadenz/pkg"
)

func NewRandom() (deadenz.Character, error) {
	return deadenz.Character{
		Type:       deadenz.CharacterType(1),
		Name:       "Test Character",
		Multiplier: 1,
	}, nil
}

func Load(b []byte) ([]deadenz.Character, error) {
	type basicCharacter struct {
		Type int    `json:"type"`
		Name string `json:"name"`
		Mult int    `json:"multiplier"`
	}

	var loaded []basicCharacter

	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, err
	}

	chars := []deadenz.Character{}

	for _, l := range loaded {
		chars = append(chars, deadenz.Character{
			Type:       deadenz.CharacterType(l.Type),
			Name:       l.Name,
			Multiplier: uint8(l.Mult),
		})
	}

	return chars, nil
}
