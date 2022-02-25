package core

type CharacterType int

const (
	MagicianCharacterType CharacterType = iota
)

type baseCharacter struct {
	name       string
	multiplier uint8
}

type Character struct {
	baseCharacter
	ctype CharacterType
}

func NewCharacter(ct CharacterType) Character {
	n, m := resolveNameAndMult(ct)

	return Character{
		ctype: ct,
		baseCharacter: baseCharacter{
			name:       n,
			multiplier: m}}
}

func (c *Character) Name() string {
	return c.name
}

func (c *Character) Multiplier() uint8 {
	return c.multiplier
}

func resolveNameAndMult(ct CharacterType) (string, uint8) {
	switch ct {
	case MagicianCharacterType:
		return "Magician", 1
	default:
		return "Useless Character", 0
	}
}
