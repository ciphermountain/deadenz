package core

type CharacterType int

const (
	MagicianCharacterType CharacterType = iota
)

type Character struct {
	type CharacterType
	name string
	multiplier uint8
}

func NewCharacter(ct CharacterType) *Character {
	n, m := resolveNameAndMult(ct)

	return &Character{
		type: ct,
		name: n,
		multiplier: m}
}

func (c *Character) Name() string {
	return c.name
}

func (c *Character) Multiplier() uint8 {
	return c.multiplier
}

func resolveNameAndMult(ct CharacterType) (string, uint8) {
	switch ct {
		case: MagicianCharacterType:
			return "Magician", 1
		default:
			return "Useless Character", 0
	}
}
