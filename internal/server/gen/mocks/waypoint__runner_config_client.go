// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	gen "github.com/hashicorp/waypoint/internal/server/gen"
	metadata "google.golang.org/grpc/metadata"

	mock "github.com/stretchr/testify/mock"
)

// Waypoint_RunnerConfigClient is an autogenerated mock type for the Waypoint_RunnerConfigClient type
type Waypoint_RunnerConfigClient struct {
	mock.Mock
}

// CloseSend provides a mock function with given fields:
func (_m *Waypoint_RunnerConfigClient) CloseSend() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Context provides a mock function with given fields:
func (_m *Waypoint_RunnerConfigClient) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// Header provides a mock function with given fields:
func (_m *Waypoint_RunnerConfigClient) Header() (metadata.MD, error) {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Recv provides a mock function with given fields:
func (_m *Waypoint_RunnerConfigClient) Recv() (*gen.RunnerConfigResponse, error) {
	ret := _m.Called()

	var r0 *gen.RunnerConfigResponse
	if rf, ok := ret.Get(0).(func() *gen.RunnerConfigResponse); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gen.RunnerConfigResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RecvMsg provides a mock function with given fields: m
func (_m *Waypoint_RunnerConfigClient) RecvMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Send provides a mock function with given fields: _a0
func (_m *Waypoint_RunnerConfigClient) Send(_a0 *gen.RunnerConfigRequest) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*gen.RunnerConfigRequest) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMsg provides a mock function with given fields: m
func (_m *Waypoint_RunnerConfigClient) SendMsg(m interface{}) error {
	ret := _m.Called(m)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(m)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Trailer provides a mock function with given fields:
func (_m *Waypoint_RunnerConfigClient) Trailer() metadata.MD {
	ret := _m.Called()

	var r0 metadata.MD
	if rf, ok := ret.Get(0).(func() metadata.MD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(metadata.MD)
		}
	}

	return r0
}
