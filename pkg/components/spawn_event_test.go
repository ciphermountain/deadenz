package components_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ciphermountain/deadenz/pkg/components"
)

func TestCharacterSpawnEvent_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	encoded := `{"type": "spawnin-event", "character": {"id": 1}}`

	var evt components.CharacterSpawnEvent

	require.NoError(t, json.Unmarshal([]byte(encoded), &evt))
}
