package deadenz_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	deadenz "github.com/ciphermountain/deadenz/pkg"
	"github.com/ciphermountain/deadenz/pkg/components"
	"github.com/ciphermountain/deadenz/pkg/middleware"
	middlewaremocks "github.com/ciphermountain/deadenz/pkg/middleware/mocks"
	"github.com/ciphermountain/deadenz/pkg/mocks"
	"github.com/ciphermountain/deadenz/pkg/opts"
)

func TestRunActionCommand_Traps(t *testing.T) {
	t.Parallel()

	trapMessage := "test trap message"
	profile := &components.Profile{
		Active:        &components.Character{},
		BackpackLimit: 10,
		Backpack:      []components.ItemType{1, 2, 3},
	}

	mTraps := new(middlewaremocks.MockTrapProvider)
	mTraps.EXPECT().TripRandom(profile, mock.AnythingOfType("opts.Option")).Return(components.Trap{Message: trapMessage}, nil)

	// test that a combination of the trap tripper middleware
	// with the walk death event and death active item middleware
	// that tripping traps functions correctly
	result, err := deadenz.RunActionCommand(
		deadenz.WalkCommandType,
		profile,
		nil,
		deadenz.Config{},
		[]deadenz.PreRunFunc{
			middleware.PreRunTrapTripper(mTraps, 1.0), // always trip
		},
		[]deadenz.PostRunFunc{
			middleware.WalkDeathEventMiddleware(),
			middleware.DeathActiveItemMiddleware(nil),
		},
		opts.WithLanguage("en"),
	)

	require.NoError(t, err)
	require.NotNil(t, result.Profile)

	assert.Equal(t, deadenz.SpawninCommandType, result.DefaultCmd)
	assert.Nil(t, result.Profile.Active)
	assert.Len(t, result.Profile.Backpack, 0)

	require.Len(t, result.Events, 1)
	assert.Equal(t, trapMessage, result.Events[0].String())
}

func TestRunActionCommand_Die(t *testing.T) {
	t.Parallel()

	profile := &components.Profile{
		Active:        &components.Character{Multiplier: 1},
		BackpackLimit: 10,
		Backpack:      []components.ItemType{1, 2, 3},
	}

	encounters := []components.EncounterEvent{components.NewEncounterEvent("a wolf")}
	actions := []components.ActionEvent{components.NewActionEvent("an action")}
	liveEvts := []components.LiveMutationEvent{}
	dieEvts := []components.DieMutationEvent{components.NewDieMutationEvent("a death", false)}

	mLoader := new(mocks.MockLoader)

	// encounter event
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.EncounterEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(encounters))

		return nil
	})

	// action event
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.ActionEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(actions))

		return nil
	})

	// live events
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.LiveMutationEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(liveEvts))

		return nil
	})

	// die events
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.DieMutationEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(dieEvts))

		return nil
	})

	// test that a combination of the trap tripper middleware
	// with the walk death event and death active item middleware
	// that tripping traps functions correctly
	result, err := deadenz.RunActionCommand(
		deadenz.WalkCommandType,
		profile,
		mLoader,
		deadenz.Config{
			DeathRate: 1.0,
		},
		[]deadenz.PreRunFunc{},
		[]deadenz.PostRunFunc{
			middleware.WalkDeathEventMiddleware(),
			middleware.DeathActiveItemMiddleware(nil),
		},
		opts.WithLanguage("en"),
	)

	require.NoError(t, err)
	require.NotNil(t, result.Profile)

	assert.Equal(t, deadenz.SpawninCommandType, result.DefaultCmd, "expect spawnin as next command")
	assert.Nil(t, result.Profile.Active)
	assert.Len(t, result.Profile.Backpack, 0)

	require.Len(t, result.Events, 5)

	assert.Equal(t, "you encounter a wolf", result.Events[0].String())
	assert.Equal(t, "an action", result.Events[1].String())
	assert.Equal(t, "a death", result.Events[2].String())
	assert.Equal(t, "you earned 1 xp", result.Events[3].String())
	assert.Equal(t, "you earned 3 tokens", result.Events[4].String())
}

func TestRunActionCommand_DieMultiverse(t *testing.T) {
	t.Parallel()

	profile := &components.Profile{
		Active:        &components.Character{Multiplier: 1},
		BackpackLimit: 10,
		Backpack:      []components.ItemType{1, 2, 3},
	}

	encounters := []components.EncounterEvent{components.NewEncounterEvent("a wolf")}
	actions := []components.ActionEvent{components.NewActionEvent("an action")}
	liveEvts := []components.LiveMutationEvent{}
	dieEvts := []components.DieMutationEvent{components.NewDieMutationEvent("a death", true)}

	mLoader := new(mocks.MockLoader)
	mEvents := new(middlewaremocks.MockEventPublisher)

	// encounter event
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.EncounterEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(encounters))

		return nil
	})

	// action event
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.ActionEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(actions))

		return nil
	})

	// live events
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.LiveMutationEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(liveEvts))

		return nil
	})

	// die events
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.DieMutationEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(dieEvts))

		return nil
	})

	mEvents.EXPECT().PublishGameEvent(mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// test that a combination of the trap tripper middleware
	// with the walk death event and death active item middleware
	// that tripping traps functions correctly
	result, err := deadenz.RunActionCommand(
		deadenz.WalkCommandType,
		profile,
		mLoader,
		deadenz.Config{
			DeathRate: 1.0,
		},
		[]deadenz.PreRunFunc{},
		[]deadenz.PostRunFunc{
			middleware.PublishEventsToMultiverse(mEvents),
			middleware.WalkDeathEventMiddleware(),
			middleware.DeathActiveItemMiddleware(nil),
		},
		opts.WithLanguage("en"),
	)

	require.NoError(t, err)
	require.NotNil(t, result.Profile)

	assert.Equal(t, deadenz.SpawninCommandType, result.DefaultCmd, "expect spawnin as next command")
	assert.Nil(t, result.Profile.Active)
	assert.Len(t, result.Profile.Backpack, 0)

	require.Len(t, result.Events, 5)

	assert.Equal(t, "you encounter a wolf", result.Events[0].String())
	assert.Equal(t, "an action", result.Events[1].String())
	assert.Equal(t, "a death", result.Events[2].String())
	assert.Equal(t, "you earned 1 xp", result.Events[3].String())
	assert.Equal(t, "you earned 3 tokens", result.Events[4].String())
}

func TestRunActionCommand_Live(t *testing.T) {
	t.Parallel()

	profile := &components.Profile{
		Active:        &components.Character{Multiplier: 1},
		BackpackLimit: 10,
		Backpack:      []components.ItemType{1, 2, 3},
	}

	encounters := []components.EncounterEvent{components.NewEncounterEvent("a wolf")}
	actions := []components.ActionEvent{components.NewActionEvent("an action")}
	liveEvts := []components.LiveMutationEvent{components.NewLiveMutationEvent("living", nil)}
	dieEvts := []components.DieMutationEvent{}

	mLoader := new(mocks.MockLoader)

	// encounter event
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.EncounterEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(encounters))

		return nil
	})

	// action event
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.ActionEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(actions))

		return nil
	})

	// live events
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.LiveMutationEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(liveEvts))

		return nil
	})

	// die events
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.DieMutationEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(dieEvts))

		return nil
	})

	// test that a combination of the trap tripper middleware
	// with the walk death event and death active item middleware
	// that tripping traps functions correctly
	result, err := deadenz.RunActionCommand(
		deadenz.WalkCommandType,
		profile,
		mLoader,
		deadenz.Config{
			DeathRate: 0.0,
		},
		[]deadenz.PreRunFunc{},
		[]deadenz.PostRunFunc{
			middleware.WalkDeathEventMiddleware(),
			middleware.DeathActiveItemMiddleware(nil),
		},
		opts.WithLanguage("en"),
	)

	require.NoError(t, err)
	require.NotNil(t, result.Profile)

	assert.Equal(t, deadenz.WalkCommandType, result.DefaultCmd, "expect walk as next command")
	assert.NotNil(t, result.Profile.Active)
	assert.Len(t, result.Profile.Backpack, 3)

	require.Len(t, result.Events, 5)

	assert.Equal(t, "you encounter a wolf", result.Events[0].String())
	assert.Equal(t, "an action", result.Events[1].String())
	assert.Equal(t, "living", result.Events[2].String())
	assert.Equal(t, "you earned 1 xp", result.Events[3].String())
	assert.Equal(t, "you earned 3 tokens", result.Events[4].String())
}

func TestRunActionCommand_LiveConvertCharacter(t *testing.T) {
	t.Parallel()

	characters := []components.Character{
		{Type: components.CharacterType(1), Name: "A", Multiplier: 1},
		{Type: components.CharacterType(2), Name: "B", Multiplier: 2},
	}

	profile := &components.Profile{
		Active:        &characters[0],
		BackpackLimit: 10,
		Backpack:      []components.ItemType{1, 2, 3},
	}

	encounters := []components.EncounterEvent{components.NewEncounterEvent("a wolf")}
	actions := []components.ActionEvent{components.NewActionEvent("an action")}
	liveEvts := []components.LiveMutationEvent{components.NewLiveMutationEvent("turn into", &characters[1].Type)}
	dieEvts := []components.DieMutationEvent{}

	mLoader := new(mocks.MockLoader)

	// encounter event
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.EncounterEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(encounters))

		return nil
	})

	// action event
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.ActionEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(actions))

		return nil
	})

	// live events
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.LiveMutationEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(liveEvts))

		return nil
	})

	// die events
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.DieMutationEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(dieEvts))

		return nil
	})

	// characters
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.Character{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(characters))

		return nil
	})

	// test that a combination of the trap tripper middleware
	// with the walk death event and death active item middleware
	// that tripping traps functions correctly
	result, err := deadenz.RunActionCommand(
		deadenz.WalkCommandType,
		profile,
		mLoader,
		deadenz.Config{
			DeathRate: 0.0,
		},
		[]deadenz.PreRunFunc{},
		[]deadenz.PostRunFunc{
			middleware.WalkDeathEventMiddleware(),
			middleware.DeathActiveItemMiddleware(nil),
		},
		opts.WithLanguage("en"),
	)

	require.NoError(t, err)
	require.NotNil(t, result.Profile)

	assert.Equal(t, deadenz.WalkCommandType, result.DefaultCmd, "expect walk as next command")
	assert.NotNil(t, result.Profile.Active)
	assert.Equal(t, characters[1], *result.Profile.Active)
	assert.Len(t, result.Profile.Backpack, 3)

	require.Len(t, result.Events, 5)

	assert.Equal(t, "you encounter a wolf", result.Events[0].String())
	assert.Equal(t, "an action", result.Events[1].String())
	assert.Equal(t, "turn into", result.Events[2].String())
	assert.Equal(t, "you earned 2 xp", result.Events[3].String())
	assert.Equal(t, "you earned 6 tokens", result.Events[4].String())
}

func TestRunActionCommand_Find(t *testing.T) {
	t.Parallel()

	profile := &components.Profile{
		Active:        &components.Character{Multiplier: 1},
		BackpackLimit: 10,
		Backpack:      []components.ItemType{1, 2, 3},
	}

	finds := []components.Item{{Name: "item", Findable: true}}
	decisions := []components.ItemDecisionEvent{components.NewItemDecisionEvent("keep")}

	mLoader := new(mocks.MockLoader)

	// find event
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.Item{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(finds))

		return nil
	})

	// decision event
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.ItemDecisionEvent{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(decisions))

		return nil
	})

	// test that a combination of the trap tripper middleware
	// with the walk death event and death active item middleware
	// that tripping traps functions correctly
	result, err := deadenz.RunActionCommand(
		deadenz.WalkCommandType,
		profile,
		mLoader,
		deadenz.Config{
			ItemFindRate: 1.0,
		},
		[]deadenz.PreRunFunc{},
		[]deadenz.PostRunFunc{
			middleware.WalkDeathEventMiddleware(),
			middleware.DeathActiveItemMiddleware(nil),
		},
		opts.WithLanguage("en"),
	)

	require.NoError(t, err)
	require.NotNil(t, result.Profile)

	assert.Equal(t, deadenz.WalkCommandType, result.DefaultCmd, "expect walk as next command")
	assert.NotNil(t, result.Profile.Active)
	assert.Len(t, result.Profile.Backpack, 3)

	require.Len(t, result.Events, 4)

	assert.Equal(t, "you find item", result.Events[0].String())
	assert.Equal(t, "keep", result.Events[1].String())
	assert.Equal(t, "you earned 1 xp", result.Events[2].String())
	assert.Equal(t, "you earned 3 tokens", result.Events[3].String())
}

func TestRunActionCommand_Spawnin(t *testing.T) {
	t.Parallel()

	profile := &components.Profile{
		Active:        nil,
		BackpackLimit: 10,
		Backpack:      []components.ItemType{},
	}

	characters := []components.Character{{Type: components.CharacterType(42), Name: "item", Multiplier: 8}}
	mLoader := new(mocks.MockLoader)

	// character load
	mLoader.EXPECT().Load(mock.MatchedBy(func(val interface{}) bool {
		return reflect.TypeOf(val) == reflect.TypeOf(&[]components.Character{})
	}), mock.AnythingOfType("opts.Option")).RunAndReturn(func(i any, _ ...opts.Option) error {
		reflect.Indirect(reflect.ValueOf(i)).Set(reflect.ValueOf(characters))

		return nil
	})

	// test that a combination of the trap tripper middleware
	// with the walk death event and death active item middleware
	// that tripping traps functions correctly
	result, err := deadenz.RunActionCommand(
		deadenz.SpawninCommandType,
		profile,
		mLoader,
		deadenz.Config{
			ItemFindRate: 1.0,
		},
		[]deadenz.PreRunFunc{},
		[]deadenz.PostRunFunc{
			middleware.WalkDeathEventMiddleware(),
			middleware.DeathActiveItemMiddleware(nil),
		},
		opts.WithLanguage("en"),
	)

	require.NoError(t, err)
	require.NotNil(t, result.Profile)

	assert.Equal(t, deadenz.WalkCommandType, result.DefaultCmd, "expect walk as next command")
	assert.NotNil(t, result.Profile.Active)
	assert.Len(t, result.Profile.Backpack, 0)

	require.Len(t, result.Events, 2)

	assert.Equal(t, "you spawned in as a item", result.Events[0].String())
	assert.Equal(t, "you earned 8 xp", result.Events[1].String())
}
