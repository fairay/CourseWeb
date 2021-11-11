package models_test

import (
	//"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects/dbuilder"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Create recipe - successful operation
*/
func TestCreateRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockAcc := new(mocks.AccountsRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)
	obj := dbuilder.RecipeMother{}.Obj0()
	author := dbuilder.AccountMother{}.Obj0()
	
	mockAcc.On("Find", obj.Author).Return(author, nil).Once()
	mockRec.On("Create", obj).Return(nil).Once()

	err := allM.Recipes.AddRecipe(obj)

	assert.Nil(t, err, "Create recipe have unexpected error")
	mockRec.AssertExpectations(t)
	mockAcc.AssertExpectations(t)
}

/*
Add grade - successful operation
*/
func TestAddGrade(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockAcc := new(mocks.AccountsRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	obj := dbuilder.RecipeMother{}.Obj0()
	acc := dbuilder.AccountMother{}.Obj0()
	
	mockRec.On("FindById", obj.Id).Return(obj, nil).Once()
	mockAcc.On("Find", acc.Login).Return(acc, nil).Once()
	mockRec.On("AddGrade", obj.Id, acc.Login).Return(nil).Once()

	err := allM.Recipes.AddGrade(obj.Id, acc.Login)

	assert.Nil(t, err, "Add grade has unexpected error")
	mockRec.AssertExpectations(t)
	mockAcc.AssertExpectations(t)
}

/*
Delete category - successful operation (recipe/category exist)
*/
/*func TestDelCategoryRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockCat := new(mocks.CategoriesRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Category = models.NewCategory(mockCat, allM)
	
	objCat := dbuilder.CategoryMother{}.Obj0()
	objRcp := dbuilder.RecipeMother{}.Obj0()

	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()

	mockCat.On("Get", objCat.Title).Return(objCat, nil).Once()

	mockCat.On("DelFromRecipe", objRcp.Id, objCat.Title).Return(nil).Once()

	err := allM.Category.DelFromRecipe(objRcp.Id, objCat)

	assert.Nil(t, err, "Deletion a category from recipe has unexpected error")
	mockRec.AssertExpectations(t)
	mockCat.AssertExpectations(t)
}*/