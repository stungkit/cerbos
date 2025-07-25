// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify
// Copyright 2021-2025 Zenauth Ltd.
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	"context"
	"io"

	"github.com/cerbos/cerbos/api/genpb/cerbos/audit/v1"
	"github.com/cerbos/cerbos/api/genpb/cerbos/response/v1"
	"github.com/cerbos/cerbos/internal/storage"
	mock "github.com/stretchr/testify/mock"
)

// NewStore creates a new instance of Store. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Store {
	mock := &Store{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

type Store_Expecter struct {
	mock *mock.Mock
}

func (_m *Store) EXPECT() *Store_Expecter {
	return &Store_Expecter{mock: &_m.Mock}
}

// Driver provides a mock function for the type Store
func (_mock *Store) Driver() string {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Driver")
	}

	var r0 string
	if returnFunc, ok := ret.Get(0).(func() string); ok {
		r0 = returnFunc()
	} else {
		r0 = ret.Get(0).(string)
	}
	return r0
}

// Store_Driver_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Driver'
type Store_Driver_Call struct {
	*mock.Call
}

// Driver is a helper method to define mock.On call
func (_e *Store_Expecter) Driver() *Store_Driver_Call {
	return &Store_Driver_Call{Call: _e.mock.On("Driver")}
}

func (_c *Store_Driver_Call) Run(run func()) *Store_Driver_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Store_Driver_Call) Return(s string) *Store_Driver_Call {
	_c.Call.Return(s)
	return _c
}

func (_c *Store_Driver_Call) RunAndReturn(run func() string) *Store_Driver_Call {
	_c.Call.Return(run)
	return _c
}

// InspectPolicies provides a mock function for the type Store
func (_mock *Store) InspectPolicies(context1 context.Context, listPolicyIDsParams storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error) {
	ret := _mock.Called(context1, listPolicyIDsParams)

	if len(ret) == 0 {
		panic("no return value specified for InspectPolicies")
	}

	var r0 map[string]*responsev1.InspectPoliciesResponse_Result
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error)); ok {
		return returnFunc(context1, listPolicyIDsParams)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, storage.ListPolicyIDsParams) map[string]*responsev1.InspectPoliciesResponse_Result); ok {
		r0 = returnFunc(context1, listPolicyIDsParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*responsev1.InspectPoliciesResponse_Result)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, storage.ListPolicyIDsParams) error); ok {
		r1 = returnFunc(context1, listPolicyIDsParams)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Store_InspectPolicies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InspectPolicies'
type Store_InspectPolicies_Call struct {
	*mock.Call
}

// InspectPolicies is a helper method to define mock.On call
//   - context1 context.Context
//   - listPolicyIDsParams storage.ListPolicyIDsParams
func (_e *Store_Expecter) InspectPolicies(context1 interface{}, listPolicyIDsParams interface{}) *Store_InspectPolicies_Call {
	return &Store_InspectPolicies_Call{Call: _e.mock.On("InspectPolicies", context1, listPolicyIDsParams)}
}

func (_c *Store_InspectPolicies_Call) Run(run func(context1 context.Context, listPolicyIDsParams storage.ListPolicyIDsParams)) *Store_InspectPolicies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 storage.ListPolicyIDsParams
		if args[1] != nil {
			arg1 = args[1].(storage.ListPolicyIDsParams)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *Store_InspectPolicies_Call) Return(stringToInspectPoliciesResponse_Result map[string]*responsev1.InspectPoliciesResponse_Result, err error) *Store_InspectPolicies_Call {
	_c.Call.Return(stringToInspectPoliciesResponse_Result, err)
	return _c
}

func (_c *Store_InspectPolicies_Call) RunAndReturn(run func(context1 context.Context, listPolicyIDsParams storage.ListPolicyIDsParams) (map[string]*responsev1.InspectPoliciesResponse_Result, error)) *Store_InspectPolicies_Call {
	_c.Call.Return(run)
	return _c
}

// ListPolicyIDs provides a mock function for the type Store
func (_mock *Store) ListPolicyIDs(context1 context.Context, listPolicyIDsParams storage.ListPolicyIDsParams) ([]string, error) {
	ret := _mock.Called(context1, listPolicyIDsParams)

	if len(ret) == 0 {
		panic("no return value specified for ListPolicyIDs")
	}

	var r0 []string
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, storage.ListPolicyIDsParams) ([]string, error)); ok {
		return returnFunc(context1, listPolicyIDsParams)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, storage.ListPolicyIDsParams) []string); ok {
		r0 = returnFunc(context1, listPolicyIDsParams)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, storage.ListPolicyIDsParams) error); ok {
		r1 = returnFunc(context1, listPolicyIDsParams)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Store_ListPolicyIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListPolicyIDs'
type Store_ListPolicyIDs_Call struct {
	*mock.Call
}

// ListPolicyIDs is a helper method to define mock.On call
//   - context1 context.Context
//   - listPolicyIDsParams storage.ListPolicyIDsParams
func (_e *Store_Expecter) ListPolicyIDs(context1 interface{}, listPolicyIDsParams interface{}) *Store_ListPolicyIDs_Call {
	return &Store_ListPolicyIDs_Call{Call: _e.mock.On("ListPolicyIDs", context1, listPolicyIDsParams)}
}

func (_c *Store_ListPolicyIDs_Call) Run(run func(context1 context.Context, listPolicyIDsParams storage.ListPolicyIDsParams)) *Store_ListPolicyIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 storage.ListPolicyIDsParams
		if args[1] != nil {
			arg1 = args[1].(storage.ListPolicyIDsParams)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *Store_ListPolicyIDs_Call) Return(strings []string, err error) *Store_ListPolicyIDs_Call {
	_c.Call.Return(strings, err)
	return _c
}

func (_c *Store_ListPolicyIDs_Call) RunAndReturn(run func(context1 context.Context, listPolicyIDsParams storage.ListPolicyIDsParams) ([]string, error)) *Store_ListPolicyIDs_Call {
	_c.Call.Return(run)
	return _c
}

// ListSchemaIDs provides a mock function for the type Store
func (_mock *Store) ListSchemaIDs(context1 context.Context) ([]string, error) {
	ret := _mock.Called(context1)

	if len(ret) == 0 {
		panic("no return value specified for ListSchemaIDs")
	}

	var r0 []string
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context) ([]string, error)); ok {
		return returnFunc(context1)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context) []string); ok {
		r0 = returnFunc(context1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = returnFunc(context1)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Store_ListSchemaIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSchemaIDs'
type Store_ListSchemaIDs_Call struct {
	*mock.Call
}

// ListSchemaIDs is a helper method to define mock.On call
//   - context1 context.Context
func (_e *Store_Expecter) ListSchemaIDs(context1 interface{}) *Store_ListSchemaIDs_Call {
	return &Store_ListSchemaIDs_Call{Call: _e.mock.On("ListSchemaIDs", context1)}
}

func (_c *Store_ListSchemaIDs_Call) Run(run func(context1 context.Context)) *Store_ListSchemaIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		run(
			arg0,
		)
	})
	return _c
}

func (_c *Store_ListSchemaIDs_Call) Return(strings []string, err error) *Store_ListSchemaIDs_Call {
	_c.Call.Return(strings, err)
	return _c
}

func (_c *Store_ListSchemaIDs_Call) RunAndReturn(run func(context1 context.Context) ([]string, error)) *Store_ListSchemaIDs_Call {
	_c.Call.Return(run)
	return _c
}

// LoadSchema provides a mock function for the type Store
func (_mock *Store) LoadSchema(context1 context.Context, s string) (io.ReadCloser, error) {
	ret := _mock.Called(context1, s)

	if len(ret) == 0 {
		panic("no return value specified for LoadSchema")
	}

	var r0 io.ReadCloser
	var r1 error
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) (io.ReadCloser, error)); ok {
		return returnFunc(context1, s)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, string) io.ReadCloser); ok {
		r0 = returnFunc(context1, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = returnFunc(context1, s)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// Store_LoadSchema_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadSchema'
type Store_LoadSchema_Call struct {
	*mock.Call
}

// LoadSchema is a helper method to define mock.On call
//   - context1 context.Context
//   - s string
func (_e *Store_Expecter) LoadSchema(context1 interface{}, s interface{}) *Store_LoadSchema_Call {
	return &Store_LoadSchema_Call{Call: _e.mock.On("LoadSchema", context1, s)}
}

func (_c *Store_LoadSchema_Call) Run(run func(context1 context.Context, s string)) *Store_LoadSchema_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 string
		if args[1] != nil {
			arg1 = args[1].(string)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *Store_LoadSchema_Call) Return(readCloser io.ReadCloser, err error) *Store_LoadSchema_Call {
	_c.Call.Return(readCloser, err)
	return _c
}

func (_c *Store_LoadSchema_Call) RunAndReturn(run func(context1 context.Context, s string) (io.ReadCloser, error)) *Store_LoadSchema_Call {
	_c.Call.Return(run)
	return _c
}

// Source provides a mock function for the type Store
func (_mock *Store) Source() *auditv1.PolicySource {
	ret := _mock.Called()

	if len(ret) == 0 {
		panic("no return value specified for Source")
	}

	var r0 *auditv1.PolicySource
	if returnFunc, ok := ret.Get(0).(func() *auditv1.PolicySource); ok {
		r0 = returnFunc()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*auditv1.PolicySource)
		}
	}
	return r0
}

// Store_Source_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Source'
type Store_Source_Call struct {
	*mock.Call
}

// Source is a helper method to define mock.On call
func (_e *Store_Expecter) Source() *Store_Source_Call {
	return &Store_Source_Call{Call: _e.mock.On("Source")}
}

func (_c *Store_Source_Call) Run(run func()) *Store_Source_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Store_Source_Call) Return(policySource *auditv1.PolicySource) *Store_Source_Call {
	_c.Call.Return(policySource)
	return _c
}

func (_c *Store_Source_Call) RunAndReturn(run func() *auditv1.PolicySource) *Store_Source_Call {
	_c.Call.Return(run)
	return _c
}
