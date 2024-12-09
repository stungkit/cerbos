// Copyright 2021-2024 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	bundle "github.com/cerbos/cloud-api/bundle"
	mock "github.com/stretchr/testify/mock"
)

// WatchHandle is an autogenerated mock type for the WatchHandle type
type WatchHandle struct {
	mock.Mock
}

type WatchHandle_Expecter struct {
	mock *mock.Mock
}

func (_m *WatchHandle) EXPECT() *WatchHandle_Expecter {
	return &WatchHandle_Expecter{mock: &_m.Mock}
}

// ActiveBundleChanged provides a mock function with given fields: _a0
func (_m *WatchHandle) ActiveBundleChanged(_a0 string) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for ActiveBundleChanged")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WatchHandle_ActiveBundleChanged_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ActiveBundleChanged'
type WatchHandle_ActiveBundleChanged_Call struct {
	*mock.Call
}

// ActiveBundleChanged is a helper method to define mock.On call
//   - _a0 string
func (_e *WatchHandle_Expecter) ActiveBundleChanged(_a0 interface{}) *WatchHandle_ActiveBundleChanged_Call {
	return &WatchHandle_ActiveBundleChanged_Call{Call: _e.mock.On("ActiveBundleChanged", _a0)}
}

func (_c *WatchHandle_ActiveBundleChanged_Call) Run(run func(_a0 string)) *WatchHandle_ActiveBundleChanged_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *WatchHandle_ActiveBundleChanged_Call) Return(_a0 error) *WatchHandle_ActiveBundleChanged_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WatchHandle_ActiveBundleChanged_Call) RunAndReturn(run func(string) error) *WatchHandle_ActiveBundleChanged_Call {
	_c.Call.Return(run)
	return _c
}

// Errors provides a mock function with no fields
func (_m *WatchHandle) Errors() <-chan error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Errors")
	}

	var r0 <-chan error
	if rf, ok := ret.Get(0).(func() <-chan error); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan error)
		}
	}

	return r0
}

// WatchHandle_Errors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Errors'
type WatchHandle_Errors_Call struct {
	*mock.Call
}

// Errors is a helper method to define mock.On call
func (_e *WatchHandle_Expecter) Errors() *WatchHandle_Errors_Call {
	return &WatchHandle_Errors_Call{Call: _e.mock.On("Errors")}
}

func (_c *WatchHandle_Errors_Call) Run(run func()) *WatchHandle_Errors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WatchHandle_Errors_Call) Return(_a0 <-chan error) *WatchHandle_Errors_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WatchHandle_Errors_Call) RunAndReturn(run func() <-chan error) *WatchHandle_Errors_Call {
	_c.Call.Return(run)
	return _c
}

// ServerEvents provides a mock function with no fields
func (_m *WatchHandle) ServerEvents() <-chan bundle.ServerEvent {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ServerEvents")
	}

	var r0 <-chan bundle.ServerEvent
	if rf, ok := ret.Get(0).(func() <-chan bundle.ServerEvent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan bundle.ServerEvent)
		}
	}

	return r0
}

// WatchHandle_ServerEvents_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ServerEvents'
type WatchHandle_ServerEvents_Call struct {
	*mock.Call
}

// ServerEvents is a helper method to define mock.On call
func (_e *WatchHandle_Expecter) ServerEvents() *WatchHandle_ServerEvents_Call {
	return &WatchHandle_ServerEvents_Call{Call: _e.mock.On("ServerEvents")}
}

func (_c *WatchHandle_ServerEvents_Call) Run(run func()) *WatchHandle_ServerEvents_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *WatchHandle_ServerEvents_Call) Return(_a0 <-chan bundle.ServerEvent) *WatchHandle_ServerEvents_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *WatchHandle_ServerEvents_Call) RunAndReturn(run func() <-chan bundle.ServerEvent) *WatchHandle_ServerEvents_Call {
	_c.Call.Return(run)
	return _c
}

// NewWatchHandle creates a new instance of WatchHandle. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewWatchHandle(t interface {
	mock.TestingT
	Cleanup(func())
}) *WatchHandle {
	mock := &WatchHandle{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
