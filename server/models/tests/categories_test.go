package models_test

import (
	"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	_ "api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCategory(t *testing.T) {
	mockRep := new(mocks.CategoriesRep)
	model := models.NewCategory(mockRep, nil)
	obj := dbuilder.CategoryMother{}.Obj0()

	mockRep.On("Get", obj.Title).Return(nil, errors.RecordNotFound).Once()
	mockRep.On("Create", obj).Return(nil).Once()

	err := model.Create(obj)

	assert.Nil(t, err, "Create have unexpected error")
	mockRep.AssertExpectations(t)
}

func TestAddCategoryRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockCat := new(mocks.CategoriesRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Category = models.NewCategory(mockCat, allM)
	obj := dbuilder.CategoryMother{}.Obj0()

	mockCat.On("Get", obj.Title).Return(nil, errors.RecordNotFound).Once()
	mockCat.On("Create", obj).Return(nil).Once()

	err := allM.Category.Create(obj)

	assert.Nil(t, err, "Create have unexpected error")
	mockCat.AssertExpectations(t)
}
