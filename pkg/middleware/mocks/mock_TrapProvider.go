// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	components "github.com/ciphermountain/deadenz/pkg/components"

	mock "github.com/stretchr/testify/mock"

	opts "github.com/ciphermountain/deadenz/pkg/opts"
)

// MockTrapProvider is an autogenerated mock type for the TrapProvider type
type MockTrapProvider struct {
	mock.Mock
}

type MockTrapProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTrapProvider) EXPECT() *MockTrapProvider_Expecter {
	return &MockTrapProvider_Expecter{mock: &_m.Mock}
}

// TripRandom provides a mock function with given fields: _a0, _a1
func (_m *MockTrapProvider) TripRandom(_a0 *components.Profile, _a1 ...opts.Option) (components.Trap, error) {
	_va := make([]interface{}, len(_a1))
	for _i := range _a1 {
		_va[_i] = _a1[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for TripRandom")
	}

	var r0 components.Trap
	var r1 error
	if rf, ok := ret.Get(0).(func(*components.Profile, ...opts.Option) (components.Trap, error)); ok {
		return rf(_a0, _a1...)
	}
	if rf, ok := ret.Get(0).(func(*components.Profile, ...opts.Option) components.Trap); ok {
		r0 = rf(_a0, _a1...)
	} else {
		r0 = ret.Get(0).(components.Trap)
	}

	if rf, ok := ret.Get(1).(func(*components.Profile, ...opts.Option) error); ok {
		r1 = rf(_a0, _a1...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockTrapProvider_TripRandom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TripRandom'
type MockTrapProvider_TripRandom_Call struct {
	*mock.Call
}

// TripRandom is a helper method to define mock.On call
//   - _a0 *components.Profile
//   - _a1 ...opts.Option
func (_e *MockTrapProvider_Expecter) TripRandom(_a0 interface{}, _a1 ...interface{}) *MockTrapProvider_TripRandom_Call {
	return &MockTrapProvider_TripRandom_Call{Call: _e.mock.On("TripRandom",
		append([]interface{}{_a0}, _a1...)...)}
}

func (_c *MockTrapProvider_TripRandom_Call) Run(run func(_a0 *components.Profile, _a1 ...opts.Option)) *MockTrapProvider_TripRandom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]opts.Option, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(opts.Option)
			}
		}
		run(args[0].(*components.Profile), variadicArgs...)
	})
	return _c
}

func (_c *MockTrapProvider_TripRandom_Call) Return(_a0 components.Trap, _a1 error) *MockTrapProvider_TripRandom_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockTrapProvider_TripRandom_Call) RunAndReturn(run func(*components.Profile, ...opts.Option) (components.Trap, error)) *MockTrapProvider_TripRandom_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTrapProvider creates a new instance of MockTrapProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTrapProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTrapProvider {
	mock := &MockTrapProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
