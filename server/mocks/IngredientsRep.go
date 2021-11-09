// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	objects "api/recipes/objects"

	mock "github.com/stretchr/testify/mock"
)

// IngredientsRep is an autogenerated mock type for the IngredientsRep type
type IngredientsRep struct {
	mock.Mock
}

// AddToRecipe provides a mock function with given fields: obj
func (_m *IngredientsRep) AddToRecipe(obj *objects.RecipeIngredient) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*objects.RecipeIngredient) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: obj
func (_m *IngredientsRep) Create(obj *objects.Ingredient) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*objects.Ingredient) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DelFromRecipe provides a mock function with given fields: idRecipe, title
func (_m *IngredientsRep) DelFromRecipe(idRecipe int, title string) error {
	ret := _m.Called(idRecipe, title)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string) error); ok {
		r0 = rf(idRecipe, title)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByRecipe provides a mock function with given fields: idRecipe
func (_m *IngredientsRep) FindByRecipe(idRecipe int) ([]objects.RecipeIngredient, error) {
	ret := _m.Called(idRecipe)

	var r0 []objects.RecipeIngredient
	if rf, ok := ret.Get(0).(func(int) []objects.RecipeIngredient); ok {
		r0 = rf(idRecipe)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]objects.RecipeIngredient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(idRecipe)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: title
func (_m *IngredientsRep) Get(title string) (*objects.Ingredient, error) {
	ret := _m.Called(title)

	var r0 *objects.Ingredient
	if rf, ok := ret.Get(0).(func(string) *objects.Ingredient); ok {
		r0 = rf(title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*objects.Ingredient)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceInRecipe provides a mock function with given fields: id_rcp, arr
func (_m *IngredientsRep) ReplaceInRecipe(id_rcp int, arr []objects.RecipeIngredient) error {
	ret := _m.Called(id_rcp, arr)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, []objects.RecipeIngredient) error); ok {
		r0 = rf(id_rcp, arr)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
