// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"
	dop "my/redditclone/pkg/interface_mongoDB"

	mock "github.com/stretchr/testify/mock"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

// PostInterface is an autogenerated mock type for the PostInterface type
type PostInterface struct {
	mock.Mock
}

// DeleteOne provides a mock function with given fields: _a0, _a1, _a2
func (_m *PostInterface) DeleteOne(_a0 context.Context, _a1 interface{}, _a2 ...*options.DeleteOptions) (dop.DeleteResultInterface, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 dop.DeleteResultInterface
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.DeleteOptions) dop.DeleteResultInterface); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dop.DeleteResultInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.DeleteOptions) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: _a0, _a1, _a2
func (_m *PostInterface) Find(_a0 context.Context, _a1 interface{}, _a2 ...*options.FindOptions) (dop.CursorInterface, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 dop.CursorInterface
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOptions) dop.CursorInterface); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dop.CursorInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.FindOptions) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertOne provides a mock function with given fields: _a0, _a1, _a2
func (_m *PostInterface) InsertOne(_a0 context.Context, _a1 interface{}, _a2 ...*options.InsertOneOptions) (dop.InsertOneResultInterface, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 dop.InsertOneResultInterface
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.InsertOneOptions) dop.InsertOneResultInterface); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dop.InsertOneResultInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.InsertOneOptions) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceOne provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *PostInterface) ReplaceOne(_a0 context.Context, _a1 interface{}, _a2 interface{}, _a3 ...*options.ReplaceOptions) (dop.UpdateResultInterface, error) {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 dop.UpdateResultInterface
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.ReplaceOptions) dop.UpdateResultInterface); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dop.UpdateResultInterface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.ReplaceOptions) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPostInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewPostInterface creates a new instance of PostInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPostInterface(t mockConstructorTestingTNewPostInterface) *PostInterface {
	mock := &PostInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
