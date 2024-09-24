package components

import (
	"encoding/json"
)

type CharacterType uint64

type Character struct {
	Type       CharacterType
	Name       string
	Multiplier uint8
}

func (c Character) MarshalJSON() ([]byte, error) {
	type jsonCharacterEncoded struct {
		ID         uint64 `json:"id"`
		Name       string `json:"string"`
		Multiplier int    `json:"multiplier"`
	}

	return json.Marshal(jsonCharacterEncoded{
		ID:         uint64(c.Type),
		Name:       c.Name,
		Multiplier: int(c.Multiplier),
	})
}

func (c *Character) UnmarshalJSON(data []byte) error {
	type jsonCharacter struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
		Mult int    `json:"multiplier"`
	}

	var loaded jsonCharacter

	if err := json.Unmarshal(data, &loaded); err != nil {
		return err
	}

	*c = Character{
		Type:       CharacterType(loaded.ID),
		Name:       loaded.Name,
		Multiplier: uint8(loaded.Mult),
	}

	return nil
}
