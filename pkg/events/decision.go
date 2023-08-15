package events

import (
	"fmt"

	"github.com/ciphermountain/deadenz/internal/util"
)

func NewRandomDecisionEvent() Event {
	decisionNames := []string{
		"add it to your backpack", // TODO: should add to the backpack
		"ignore it",
		"eat it",
		"look at it inquisitively",
		"pour water on it",
		"pretend it's a microphone and you sing",
		"mistake it for a water bottle and you drink from it",
		"burn it",
		"throw it in the rubbish bin",
		"play baseball with it",
		"feed it to your pet tiger in Oklahoma",
		"throw it at your best friend",
	}

	idx := int(util.Random(0, int64(len(decisionNames)-1)))

	return &LiveMutationEvent{value: decisionNames[idx]}
}

type DecisionEvent struct {
	value string
}

func (e DecisionEvent) String() string {
	return fmt.Sprintf("you %s", e.value)
}
