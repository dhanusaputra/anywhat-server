// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	pb "github.com/dhanusaputra/anywhat-server/api/pb"
	mock "github.com/stretchr/testify/mock"
)

// Anywhat is an autogenerated mock type for the Anywhat type
type Anywhat struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, anything
func (_m *Anywhat) Create(ctx context.Context, anything *pb.Anything) (string, error) {
	ret := _m.Called(ctx, anything)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *pb.Anything) string); ok {
		r0 = rf(ctx, anything)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.Anything) error); ok {
		r1 = rf(ctx, anything)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Anywhat) Delete(ctx context.Context, id string) (bool, error) {
	ret := _m.Called(ctx, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, id
func (_m *Anywhat) Get(ctx context.Context, id string) (*pb.Anything, error) {
	ret := _m.Called(ctx, id)

	var r0 *pb.Anything
	if rf, ok := ret.Get(0).(func(context.Context, string) *pb.Anything); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Anything)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx
func (_m *Anywhat) List(ctx context.Context) ([]*pb.Anything, error) {
	ret := _m.Called(ctx)

	var r0 []*pb.Anything
	if rf, ok := ret.Get(0).(func(context.Context) []*pb.Anything); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pb.Anything)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, anything
func (_m *Anywhat) Update(ctx context.Context, anything *pb.Anything) (bool, error) {
	ret := _m.Called(ctx, anything)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *pb.Anything) bool); ok {
		r0 = rf(ctx, anything)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pb.Anything) error); ok {
		r1 = rf(ctx, anything)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
