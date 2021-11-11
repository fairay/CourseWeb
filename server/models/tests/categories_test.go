package models_test

import (
	"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	_ "api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"api/recipes/repository"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/// CLASSIC STYLE (Stub)

/*
Get all category - sucess
*/
func TestGetAll(t *testing.T) {
	db, err := stubConnecton();
	if err != nil { panic(err); }

	objArr := dbuilder.CategoryMother{}.All()

	mockRep := repository.NewCategotiesStub(db)
	for _, obj := range objArr {
		if err := mockRep.Create(&obj); err != nil { panic(err); }
	}

	model := models.NewCategory(mockRep, nil)

	resArr := model.GetAll()
	fmt.Println(resArr)

	assert.ElementsMatch(t, resArr, objArr)
}

/// LONDON STYLE (Mock)

/*
Create recipe - successful operation
*/
func TestCreateCategory(t *testing.T) {
	mockRep := new(mocks.CategoriesRep)
	model := models.NewCategory(mockRep, nil)
	obj := dbuilder.CategoryMother{}.Obj0()

	mockRep.On("Get", obj.Title).Return(nil, errors.RecordNotFound).Once()
	mockRep.On("Create", obj).Return(nil).Once()

	err := model.Create(obj)

	assert.Nil(t, err, "Create category have unexpected error")
	mockRep.AssertExpectations(t)
}

/*
func TestAddCategoryRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockCat := new(mocks.CategoriesRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Category = models.NewCategory(mockCat, allM)
	obj := dbuilder.CategoryMother{}.Obj0()

	mockRec.On("")
	mockCat.On("Get", obj.Title).Return(nil, errors.RecordNotFound).Once()
	mockCat.On("Create", obj).Return(nil).Once()

	err := allM.Category.Create(obj)

	assert.Nil(t, err, "Create have unexpected error")
	mockCat.AssertExpectations(t)
}

func TestDelCategoryRecipe(t *testing.T) {
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
*/