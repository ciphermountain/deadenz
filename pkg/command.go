package deadenz

type CommandType int

const (
	ExitCommandType CommandType = iota
	SpawninCommandType
	WalkCommandType
	BackpackCommandType
	XPCommandType
	CurrencyCommandType
)
