// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// HandlerMessage is an autogenerated mock type for the HandlerMessage type
type HandlerMessage struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, eventMessage
func (_m *HandlerMessage) Execute(ctx context.Context, eventMessage []byte) error {
	ret := _m.Called(ctx, eventMessage)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []byte) error); ok {
		r0 = rf(ctx, eventMessage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewHandlerMessage creates a new instance of HandlerMessage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandlerMessage(t interface {
	mock.TestingT
	Cleanup(func())
}) *HandlerMessage {
	mock := &HandlerMessage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
