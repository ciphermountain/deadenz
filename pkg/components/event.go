package components

type Event interface {
	String() string
}

type EventType string

const (
	EventTypeAction       EventType = "action"
	EventTypeItemDecision EventType = "item_decision"
	EventTypeEncounter    EventType = "encounter"
	EventTypeFind         EventType = "find"
	EventTypeMutation     EventType = "mutation"
	EventTypeSpawnin      EventType = "spawnin"
)
