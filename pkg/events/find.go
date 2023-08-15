package events

import (
	"fmt"

	deadenz "github.com/ciphermountain/deadenz/pkg"
)

func NewFindEvent(item deadenz.Item) Event {
	return FindEvent{Item: item}
}

type FindEvent struct {
	Item deadenz.Item
}

func (e FindEvent) String() string {
	return fmt.Sprintf("you find %s", e.Item.Name)
}
