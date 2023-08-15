package characters

import deadenz "github.com/ciphermountain/deadenz/pkg"

func NewRandom() (deadenz.Character, error) {
	return deadenz.Character{
		Type:       deadenz.CharacterType(1),
		Name:       "Test Character",
		Multiplier: 1,
	}, nil
}
