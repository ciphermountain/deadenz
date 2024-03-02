package events

import (
	"fmt"

	"github.com/ciphermountain/deadenz/internal/util"
)

func NewRandomMutationEvent() Event {
	if util.Random(0, 100) < 30 {
		return NewRandomDieMutationEvent()
	}

	return NewRandomLiveMutationEvent()
}

func NewRandomDieMutationEvent() Event {
	waysToDie := []string{
		"die instantly",
		"die and become a ghost",
		"die of dysentery",
		"get eaten ... and die",
		"die from contemplating incongruencias in the space time continuum",
		"die from crossing your eyes to impress your friends one too many times",
		"get abducted by aliens and never return to your home planet; you’re as good as dead",
		"and all the other yous in the multiverse die", // TODO: causes all others with the same character to die
		"explode … and die",
		"implode … and die",
		"get swallowed by quicksand and die",
		"stub your toe in the dark and die",
		"die by the sting of 1000 bees",
		"find a bone in your hot dog and die",
		"experience a cuteness overload and die",
		"die from an epic rugpull",
		"die of starvation thinking of ways to die",
		"die of old age",
		"die of young age",
		"die from a medieval executioner",
		"die of measles",
		"die and the grim reaper haunts you forever",
	}

	idx := int(util.Random(0, int64(len(waysToDie)-1)))

	return &DieMutationEvent{value: waysToDie[idx]}
}

type DieMutationEvent struct {
	value string
}

func (e DieMutationEvent) String() string {
	return fmt.Sprintf("you %s", e.value)
}

func NewRandomLiveMutationEvent() Event {
	waysToLive := []string{
		"turn into a fish", // TODO: should switch character types
		"turn into a pencil",
		"turn into an inverted mermaid",
		"turn into a ball of string",
		"turn into a gummy bear",
		"survive a deadly encounter",
		"marry it and have two beautiful children Nathaniel and Supa Fly",
		"realize it is also a comic loving nerd just like you and you make a new friend",
		"are told to put your hands on the oodles of noodles and you ask chicken or beef",
		"are forced to sing along to a song you don't know",
		"get coal for Christmas",
		"feel surprised that nothing happened",
		"get an F- in geography",
		"receive an attaboy from your dear old dad",
		"discover a strange new addiction",
		"discover a strange new obsession",
		"learn a new skill",
		"get a permanent pizza stain on your upper lip",
		"are given a first class one way ticket to Albuquerque",
		"get hired to work the Night Shift at a pizzeria",
		"find that your terror farts just saved your life",
		"moon walk out of the situation",
		"capitalize on the confusion and run away safe",
	}

	idx := int(util.Random(0, int64(len(waysToLive)-1)))

	return &LiveMutationEvent{value: waysToLive[idx]}
}

type LiveMutationEvent struct {
	value string
}

func (e LiveMutationEvent) String() string {
	return fmt.Sprintf("you %s", e.value)
}
