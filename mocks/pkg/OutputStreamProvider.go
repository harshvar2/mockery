// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	pkg "github.com/vektra/mockery/v2/pkg"
)

// OutputStreamProvider is an autogenerated mock type for the OutputStreamProvider type
type OutputStreamProvider struct {
	mock.Mock
}

// GetWriter provides a mock function with given fields: _a0, _a1
func (_m *OutputStreamProvider) GetWriter(_a0 context.Context, _a1 *pkg.Interface) (io.Writer, error, pkg.Cleanup) {
	ret := _m.Called(_a0, _a1)

	var r0 io.Writer
	if rf, ok := ret.Get(0).(func(context.Context, *pkg.Interface) io.Writer); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.Writer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *pkg.Interface) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	var r2 pkg.Cleanup
	if rf, ok := ret.Get(2).(func(context.Context, *pkg.Interface) pkg.Cleanup); ok {
		r2 = rf(_a0, _a1)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(pkg.Cleanup)
		}
	}

	return r0, r1, r2
}
