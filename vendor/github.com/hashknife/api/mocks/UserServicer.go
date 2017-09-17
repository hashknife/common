package mocks

import "github.com/stretchr/testify/mock"

import "context"

type UserServicer struct {
	mock.Mock
}

// Retrieve provides a mock function with given fields: ctx, i
func (_m *UserServicer) Retrieve(ctx context.Context, i interface{}) (interface{}, error) {
	ret := _m.Called(ctx, i)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) interface{}); ok {
		r0 = rf(ctx, i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, i
func (_m *UserServicer) Update(ctx context.Context, i interface{}) (interface{}, error) {
	ret := _m.Called(ctx, i)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) interface{}); ok {
		r0 = rf(ctx, i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Disable provides a mock function with given fields: ctx, i
func (_m *UserServicer) Disable(ctx context.Context, i interface{}) (interface{}, error) {
	ret := _m.Called(ctx, i)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) interface{}); ok {
		r0 = rf(ctx, i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, i)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
