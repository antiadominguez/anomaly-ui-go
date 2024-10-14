// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	interfaces "github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	mock "github.com/stretchr/testify/mock"

	types "github.com/edgexfoundry/go-mod-messaging/v3/pkg/types"
)

// TriggerContextBuilder is an autogenerated mock type for the TriggerContextBuilder type
type TriggerContextBuilder struct {
	mock.Mock
}

// Execute provides a mock function with given fields: env
func (_m *TriggerContextBuilder) Execute(env types.MessageEnvelope) interfaces.AppFunctionContext {
	ret := _m.Called(env)

	var r0 interfaces.AppFunctionContext
	if rf, ok := ret.Get(0).(func(types.MessageEnvelope) interfaces.AppFunctionContext); ok {
		r0 = rf(env)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interfaces.AppFunctionContext)
		}
	}

	return r0
}

type mockConstructorTestingTNewTriggerContextBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewTriggerContextBuilder creates a new instance of TriggerContextBuilder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTriggerContextBuilder(t mockConstructorTestingTNewTriggerContextBuilder) *TriggerContextBuilder {
	mock := &TriggerContextBuilder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}