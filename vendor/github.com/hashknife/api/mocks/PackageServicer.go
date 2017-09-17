package mocks

import "github.com/stretchr/testify/mock"

import "context"

type PackageServicer struct {
	mock.Mock
}

// Status provides a mock function with given fields: ctx, i
func (_m *PackageServicer) Status(ctx context.Context, i interface{}) (interface{}, error) {
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

// Deliver provides a mock function with given fields: ctx, i
func (_m *PackageServicer) Deliver(ctx context.Context, i interface{}) (interface{}, error) {
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

// Accept provides a mock function with given fields: ctx, i
func (_m *PackageServicer) Accept(ctx context.Context, i interface{}) (interface{}, error) {
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

// Request provides a mock function with given fields: ctx, i
func (_m *PackageServicer) Request(ctx context.Context, i interface{}) (interface{}, error) {
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
