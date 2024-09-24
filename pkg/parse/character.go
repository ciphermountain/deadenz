package parse

import (
	"encoding/json"

	"github.com/ciphermountain/deadenz/pkg/components"
)

func CharactersFromJSON(b []byte) ([]components.Character, error) {
	var loaded []components.Character

	if err := json.Unmarshal(b, &loaded); err != nil {
		return nil, err
	}

	return loaded, nil
}
