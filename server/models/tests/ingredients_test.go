package models_test

import (
	"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects/dbuilder"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Create ingredient - successful operation
*/
func TestCreateIngredient(t *testing.T) {
	mockRep := new(mocks.IngredientsRep)
	model := models.NewIngredient(mockRep, nil)
	obj := dbuilder.IngredientMother{}.Obj0()

	mockRep.On("Get", obj.Title).Return(nil, errors.RecordNotFound).Once()
	mockRep.On("Create", obj).Return(nil).Once()

	err := model.Create(obj)

	assert.Nil(t, err, "Create ingredient have unexpected error")
	mockRep.AssertExpectations(t)
}