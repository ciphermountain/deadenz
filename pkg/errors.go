package deadenz

import "github.com/ciphermountain/deadenz/pkg/components"

type ErrTrapTripped struct {
	Trap components.Trap
}

func (e ErrTrapTripped) Error() string {
	return "trap tripped"
}

func (e ErrTrapTripped) Is(target error) bool {
	_, ok := target.(ErrTrapTripped)

	return ok
}
