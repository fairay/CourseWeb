// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	objects "api/recipes/objects"

	mock "github.com/stretchr/testify/mock"
)

// StepsRep is an autogenerated mock type for the StepsRep type
type StepsRep struct {
	mock.Mock
}

// Create provides a mock function with given fields: obj
func (_m *StepsRep) Create(obj *objects.Step) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*objects.Step) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id_rcp, step
func (_m *StepsRep) Delete(id_rcp int, step int) error {
	ret := _m.Called(id_rcp, step)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(id_rcp, step)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindStepByNum provides a mock function with given fields: id_rcp, step
func (_m *StepsRep) FindStepByNum(id_rcp int, step int) (objects.Step, error) {
	ret := _m.Called(id_rcp, step)

	var r0 objects.Step
	if rf, ok := ret.Get(0).(func(int, int) objects.Step); ok {
		r0 = rf(id_rcp, step)
	} else {
		r0 = ret.Get(0).(objects.Step)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(id_rcp, step)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindStepLast provides a mock function with given fields: id_rcp
func (_m *StepsRep) FindStepLast(id_rcp int) objects.Step {
	ret := _m.Called(id_rcp)

	var r0 objects.Step
	if rf, ok := ret.Get(0).(func(int) objects.Step); ok {
		r0 = rf(id_rcp)
	} else {
		r0 = ret.Get(0).(objects.Step)
	}

	return r0
}

// FindSteps provides a mock function with given fields: id_rcp
func (_m *StepsRep) FindSteps(id_rcp int) ([]objects.Step, error) {
	ret := _m.Called(id_rcp)

	var r0 []objects.Step
	if rf, ok := ret.Get(0).(func(int) []objects.Step); ok {
		r0 = rf(id_rcp)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]objects.Step)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id_rcp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: recipeId
func (_m *StepsRep) List(recipeId int) ([]objects.Step, error) {
	ret := _m.Called(recipeId)

	var r0 []objects.Step
	if rf, ok := ret.Get(0).(func(int) []objects.Step); ok {
		r0 = rf(recipeId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]objects.Step)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(recipeId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStep provides a mock function with given fields: id_rcp, step, obj
func (_m *StepsRep) UpdateStep(id_rcp int, step int, obj *objects.Step) error {
	ret := _m.Called(id_rcp, step, obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int, *objects.Step) error); ok {
		r0 = rf(id_rcp, step, obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}