package events

import (
	"encoding/json"

	"github.com/ciphermountain/deadenz/internal/util"
)

const DefaultDieRate = 30

func NewRandomMutationEvent(live []LiveMutationEvent, die []DieMutationEvent, diePercent int64) Event {
	if util.Random(0, 100) < diePercent {
		return die[util.Random(0, int64(len(die)-1))]
	}

	return live[util.Random(0, int64(len(live)-1))]
}

type DieMutationEvent struct {
	value string
}

func (e DieMutationEvent) String() string {
	return e.value
}

func LoadMutations(b []byte) ([]LiveMutationEvent, []DieMutationEvent, error) {
	type action struct {
		Message string `json:"message"`
		IsDeath bool   `json:"isDeath"`
	}

	var loaded []action

	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, nil, err
	}

	liveevts := []LiveMutationEvent{}
	dieEvts := []DieMutationEvent{}

	for _, l := range loaded {
		if !l.IsDeath {
			liveevts = append(liveevts, LiveMutationEvent{
				value: l.Message,
			})
		} else {
			dieEvts = append(dieEvts, DieMutationEvent{
				value: l.Message,
			})
		}
	}

	return liveevts, dieEvts, nil
}

type LiveMutationEvent struct {
	value string
}

func (e LiveMutationEvent) String() string {
	return e.value
}
