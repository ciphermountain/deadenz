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
	"github.com/ciphermountain/deadenz/pkg/opts"
)

func TestDataLoader(t *testing.T) {
	dataLoader := data.NewDataLoader()
	loadedType := reflect.TypeOf([]string{})
	loader := new(mocks.MockLoader)

	require.NoError(t, dataLoader.SetLoader(loadedType, loader, json.Unmarshal))
	loader.EXPECT().Data(mock.Anything).Return(encoded, nil)

	var output []string

	err := dataLoader.LoadCtx(context.Background(), &output)

	require.NoError(t, err)
	assert.Equal(t, strVals, output)
}

func TestDataLoader_LanguageScoped(t *testing.T) {
	dataLoader := data.NewDataLoader()
	loadedType := reflect.TypeOf([]string{})
	loaderA := new(mocks.MockLoader)
	loaderB := new(mocks.MockLoader)

	require.NoError(t, dataLoader.SetLoader(loadedType, loaderA, json.Unmarshal, opts.WithLanguage("en")))
	require.NoError(t, dataLoader.SetLoader(loadedType, loaderB, json.Unmarshal, opts.WithLanguage("es")))

	loaderA.EXPECT().Data(mock.Anything).Return(encoded, nil)
	loaderB.EXPECT().Data(mock.Anything).Return(encodedB, nil)

	var outputA []string

	require.NoError(t, dataLoader.LoadCtx(context.Background(), &outputA, opts.WithLanguage("en")))
	assert.Equal(t, strVals, outputA)

	loaderA.AssertExpectations(t)

	var outputB []string

	require.NoError(t, dataLoader.LoadCtx(context.Background(), &outputB, opts.WithLanguage("es")))
	assert.Equal(t, strValsB, outputB)

	loaderB.AssertExpectations(t)
}

func TestDataLoader_WithReloadInterval(t *testing.T) {
	dataLoader := data.NewDataLoader()

	require.NoError(t, dataLoader.Start())

	t.Cleanup(func() {
		require.NoError(t, dataLoader.Close())
	})

	loadedType := reflect.TypeOf([]string{})
	loaderA := new(mocks.MockLoader)
	loaderB := new(mocks.MockLoader)

	require.NoError(t, dataLoader.SetLoader(loadedType, loaderA, json.Unmarshal, opts.WithLanguage("en"), data.WithReloadInterval(500*time.Millisecond)))
	require.NoError(t, dataLoader.SetLoader(loadedType, loaderB, json.Unmarshal, opts.WithLanguage("es"), data.WithReloadInterval(500*time.Millisecond)))

	var loadedA atomic.Bool

	loaderA.EXPECT().Data(mock.Anything).RunAndReturn(func(ctx context.Context) ([]byte, error) {
		if loadedA.Load() {
			return encoded, nil
		} else {
			loadedA.Store(true)

			return json.Marshal(append(strVals, "zero"))
		}
	}).Twice()

	var loadedB atomic.Bool

	loaderB.EXPECT().Data(mock.Anything).RunAndReturn(func(ctx context.Context) ([]byte, error) {
		if loadedB.Load() {
			return encodedB, nil
		} else {
			loadedB.Store(true)

			return json.Marshal(append(strValsB, "zero"))
		}
	}).Twice()

	var callCount int

	assert.Eventually(t, func() bool {
		var outputA []string

		require.NoError(t, dataLoader.LoadCtx(context.Background(), &outputA, opts.WithLanguage("en")))

		var outputB []string

		require.NoError(t, dataLoader.LoadCtx(context.Background(), &outputB, opts.WithLanguage("es")))

		callCount++

		return reflect.DeepEqual(strVals, outputA) && reflect.DeepEqual(strValsB, outputB) && callCount >= 3
	}, 2*time.Second, 250*time.Millisecond)

	loaderA.AssertExpectations(t)
	loaderB.AssertExpectations(t)
}

var (
	strVals     = []string{"one", "two", "three"}
	strValsB    = []string{"four", "five", "six"}
	encoded, _  = json.Marshal(strVals)
	encodedB, _ = json.Marshal(strValsB)
)
