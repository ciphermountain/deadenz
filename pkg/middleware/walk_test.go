package middleware_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/middleware"
	"github.com/ciphermountain/deadenz/pkg/middleware/mocks"
)

var itemID components.ItemType = 42

func TestWalkLimiter(t *testing.T) {
	t.Parallel()

	mockItemProvider := mocks.NewMockItemProvider(t)
	limiter := middleware.WalkLimiter(2, mockItemProvider)

	t.Run("Failure", func(t *testing.T) {
		t.Parallel()

		t.Run("non-walk command should not return error and not modify profile", func(t *testing.T) {
			t.Parallel()

			profile := &components.Profile{}
			newProfile, err := limiter(deadenz.SpawninCommandType, profile)

			require.NoError(t, err)
			assert.Equal(t, profile, newProfile)
		})

		t.Run("nil profile should return error", func(t *testing.T) {
			t.Parallel()

			profile, err := limiter(deadenz.WalkCommandType, nil)

			require.ErrorIs(t, err, middleware.ErrNilProfile)
			assert.Nil(t, profile)
		})

		t.Run("walk too much should return error", func(t *testing.T) {
			t.Parallel()

			profile := &components.Profile{
				Limits: &components.Limits{
					LastWalk:  time.Now().Add(-1 * time.Minute),
					WalkCount: 3,
				},
			}
			newProfile, err := limiter(deadenz.WalkCommandType, profile)

			require.ErrorIs(t, err, middleware.ErrWalkTooMuch)
			assert.Equal(t, profile, newProfile)
		})
	})

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		t.Run("no walk limits set on profile should set limits to initial values", func(t *testing.T) {
			t.Parallel()

			start := time.Now()
			profile := &components.Profile{}
			newProfile, err := limiter(deadenz.WalkCommandType, profile)

			require.NoError(t, err)
			require.NotNil(t, newProfile.Limits)

			assert.Equal(t, uint64(1), newProfile.Limits.WalkCount)
			assert.GreaterOrEqual(t, newProfile.Limits.LastWalk, start)
		})

		t.Run("valid walk increases walk count by 1", func(t *testing.T) {
			t.Parallel()

			start := time.Now()
			profile := &components.Profile{
				Limits: &components.Limits{
					LastWalk:  time.Now().Add(-1 * time.Minute),
					WalkCount: 1,
				},
			}
			newProfile, err := limiter(deadenz.WalkCommandType, profile)

			require.NoError(t, err)
			require.NotNil(t, newProfile.Limits)

			assert.Equal(t, uint64(2), newProfile.Limits.WalkCount)
			assert.GreaterOrEqual(t, newProfile.Limits.LastWalk.UnixMilli(), start.UnixMilli())
		})

		t.Run("active item increases walking limit", func(t *testing.T) {
			t.Parallel()

			item := components.Item{
				Type: itemID,
				Name: "test",
				Usability: &components.Usability{
					ImprovesWalking: true,
					Efficiency:      components.DefaultSkillEfficiency,
				},
			}

			mockItemProvider.EXPECT().Item(mock.Anything).Return(&item, nil)

			start := time.Now()
			profile := &components.Profile{
				ActiveItem: &itemID,
				Limits: &components.Limits{
					LastWalk:  time.Now().Add(-1 * time.Minute),
					WalkCount: 3,
				},
			}
			newProfile, err := limiter(deadenz.WalkCommandType, profile)

			require.NoError(t, err)
			require.NotNil(t, newProfile.Limits)

			assert.Equal(t, uint64(4), newProfile.Limits.WalkCount)
			assert.GreaterOrEqual(t, newProfile.Limits.LastWalk.UnixMilli(), start.UnixMilli())
		})

		t.Run("recent walk after long wait upcounts walk limit", func(t *testing.T) {
			t.Parallel()

			start := time.Now()
			profile := &components.Profile{
				Limits: &components.Limits{
					LastWalk:  time.Now().Add(-1 * time.Hour),
					WalkCount: 3,
				},
			}
			newProfile, err := limiter(deadenz.WalkCommandType, profile)

			require.NoError(t, err)
			require.NotNil(t, newProfile.Limits)

			assert.Equal(t, uint64(2), newProfile.Limits.WalkCount)
			assert.GreaterOrEqual(t, newProfile.Limits.LastWalk.UnixMilli(), start.UnixMilli())
		})
	})
}
