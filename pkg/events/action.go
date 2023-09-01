package events

import (
	"fmt"

	"github.com/ciphermountain/deadenz/internal/util"
)

func NewRandomActionEvent() Event {
	actions := []string{
		"challenge it to a rap battle",
		"fight it",
		"date it",
		"eat it",
		"punch it",
		"offer it a sandwich",
		"run away",
		"give it scritches",
		"offer it your pants",
		"clip its fingernails",
		"challenge it to tick tack toe",
		"challenge it to a game of paintball",
		"give it a bath",
		"give it a bat",
		"live in a house it built for you",
		"offer a sacrifice to the Aztec gods",
		"do your homework",
		"burp loudly",
		"fart in terror",
		"start a dance off",
		"throw some fortune cookies",
		"fart in its general derection",
		"play dead ends for 80 hours straight",
		"become 90 minutes of someones day if you know what i mean [you play catan twice with them]",
		"you you you you you you you yooooooy yoouuuu uoy oyu uyo",
		"UWU",
		"UWU UWU UWU UWU UWU UWU UWU UWU UWU UWU UWU",
		"you say UWU and it now wants to date you",
		"you mistake it for a water bottle and you drink from it",
		"you play a chekin (che-keen) game for 365 days straight",
	}

	idx := int(util.Random(0, int64(len(actions)-1)))

	return &ActionEvent{value: actions[idx]}
}

type ActionEvent struct {
	value string
}

func (e ActionEvent) String() string {
	return fmt.Sprintf("you decide to %s", e.value)
}
