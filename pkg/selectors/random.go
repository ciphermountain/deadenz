package selectors

import "github.com/ciphermountain/deadenz/pkg/components"

func RandomCharacter() (components.Character, error) {
	return components.Character{
		Type:       components.CharacterType(1),
		Name:       "Test Character",
		Multiplier: 1,
	}, nil
}
