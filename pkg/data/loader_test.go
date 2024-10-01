package data_test

import (
	"context"
	"encoding/json"
	"reflect"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/ciphermountain/deadenz/pkg/data"
	"github.com/ciphermountain/deadenz/pkg/data/mocks"
)

func TestDataLoader(t *testing.T) {
	dataLoader := data.NewDataLoader()
	loadedType := reflect.TypeOf([]string{})
	loader := new(mocks.MockLoader)

	dataLoader.SetLoader(loadedType, loader, json.Unmarshal)
	loader.EXPECT().Data(mock.Anything).Return(encoded, nil)

	var output []string

	err := dataLoader.LoadCtx(context.Background(), &output)

	require.NoError(t, err)
	assert.Equal(t, strVals, output)
}

func TestDataLoader_WithReloadInterval(t *testing.T) {
	dataLoader := data.NewDataLoader()

	require.NoError(t, dataLoader.Start())

	t.Cleanup(func() {
		require.NoError(t, dataLoader.Close())
	})

	loadedType := reflect.TypeOf([]string{})
	loader := new(mocks.MockLoader)

	dataLoader.SetLoader(loadedType, loader, json.Unmarshal, data.WithReloadInterval(500*time.Millisecond))

	var loaded atomic.Bool

	loader.EXPECT().Data(mock.Anything).RunAndReturn(func(ctx context.Context) ([]byte, error) {
		if loaded.Load() {
			return encoded, nil
		} else {
			loaded.Store(true)

			return json.Marshal([]string{"zero", "one", "two"})
		}
	}).Twice()

	var callCount int

	assert.Eventually(t, func() bool {
		var output []string

		require.NoError(t, dataLoader.LoadCtx(context.Background(), &output))

		callCount++

		return reflect.DeepEqual(strVals, output) && callCount >= 3
	}, 2*time.Second, 250*time.Millisecond)

	loader.AssertExpectations(t)
}

var (
	strVals    = []string{"one", "two", "three"}
	encoded, _ = json.Marshal(strVals)
)
