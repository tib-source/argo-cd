// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package mocks

import (
	mock "github.com/stretchr/testify/mock"
)

// NewRbacEnforcer creates a new instance of RbacEnforcer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRbacEnforcer(t interface {
	mock.TestingT
	Cleanup(func())
}) *RbacEnforcer {
	mock := &RbacEnforcer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// RbacEnforcer is an autogenerated mock type for the RbacEnforcer type
type RbacEnforcer struct {
	mock.Mock
}

type RbacEnforcer_Expecter struct {
	mock *mock.Mock
}

func (_m *RbacEnforcer) EXPECT() *RbacEnforcer_Expecter {
	return &RbacEnforcer_Expecter{mock: &_m.Mock}
}

// EnforceErr provides a mock function for the type RbacEnforcer
func (_mock *RbacEnforcer) EnforceErr(rvals ...any) error {
	var _ca []interface{}
	_ca = append(_ca, rvals...)
	ret := _mock.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for EnforceErr")
	}

	var r0 error
	if returnFunc, ok := ret.Get(0).(func(...any) error); ok {
		r0 = returnFunc(rvals...)
	} else {
		r0 = ret.Error(0)
	}
	return r0
}

// RbacEnforcer_EnforceErr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EnforceErr'
type RbacEnforcer_EnforceErr_Call struct {
	*mock.Call
}

// EnforceErr is a helper method to define mock.On call
//   - rvals ...any
func (_e *RbacEnforcer_Expecter) EnforceErr(rvals ...interface{}) *RbacEnforcer_EnforceErr_Call {
	return &RbacEnforcer_EnforceErr_Call{Call: _e.mock.On("EnforceErr",
		append([]interface{}{}, rvals...)...)}
}

func (_c *RbacEnforcer_EnforceErr_Call) Run(run func(rvals ...any)) *RbacEnforcer_EnforceErr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 []any
		variadicArgs := make([]any, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(any)
			}
		}
		arg0 = variadicArgs
		run(
			arg0...,
		)
	})
	return _c
}

func (_c *RbacEnforcer_EnforceErr_Call) Return(err error) *RbacEnforcer_EnforceErr_Call {
	_c.Call.Return(err)
	return _c
}

func (_c *RbacEnforcer_EnforceErr_Call) RunAndReturn(run func(rvals ...any) error) *RbacEnforcer_EnforceErr_Call {
	_c.Call.Return(run)
	return _c
}
