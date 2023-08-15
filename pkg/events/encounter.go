package events

import (
	"fmt"

	"github.com/ciphermountain/deadenz/internal/util"
)

func NewRandomEncounterEvent() Event {
	encounters := []string{
		"a microwave that wants to date you",
		"Deery McDeerface",
		"a giant cancer blob",
		"a creature so hideous words cannot describe it",
		"a fish",
		"an anime zombie creature",
		"a Mayan god",
		"a creature so beautiful words cannot describe it",
		"a creapy clown",
		"sargeant Boxer Shorts",
		"Santa Claus",
		"a living sweaty gym sock",
		"a poisonous mushroom",
		"Swordy McSwordface",
		"a fairy",
		"a demon with perfect fingernails",
		"a fountain of youth",
		"a dog with rabies",
		"a bunny",
		"herobrine",
		"a scarlet macaw",
		"a baby shark",
		"a living apple",
		"a cucumber named larry",
		"shawoecapooenope the dinosaur",
		"jif peanut butter",
		"a rat",
		"a hippo",
		"a pizza monster",
		"some living sunglasses",
		"a hypocrite",
		"Skelly Jack",
		"Smelly Jack",
		"a living donut spider",
		"the tooth fairy",
		"a quantum recursion trap",
		"an outdated meme",
		"an anaconda",
		"a tax investment winged lizard that breathes fire",
		"a tax insurance dragon",
		"a juggling frog",
		"a cookie cat",
		"a French cassowary",
		"a burrito sabanero que camine a belen",
		"a savage Norseman",
		"a creature with the head of a horse and the body of a cat",
		"a horde of succulent shrimp",
		"some flesh-eating elves",
		"a werewolf",
		"the oracle",
		"a herd of baby goats",
		"a Gorn",
		"a spooky ghost",
		"a creepy doll",
		"a messy potato man",
		"an inverted mermaid",
		"a smurf",
		"a big stinky fart",
		"a tree",
		"nothing",
		"a huge deasel enjine druck",
		"a duck",
		"a big stinky diaper",
		"the man with the upside down face",
		"the little green car",
		"xcaret",
		"chlorine",
		"ERROR PAGE NOT FOUND",
		"a huge moth that eats humans",
		"a very cute doggy",
		"a TOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOT",
		"a tiny tot",
	}

	idx := int(util.Random(0, int64(len(encounters)-1)))

	return &EncounterEvent{value: encounters[idx]}
}

type EncounterEvent struct {
	value string
}

func (e EncounterEvent) String() string {
	return fmt.Sprintf("you encounter %s", e.value)
}
