package models_test

import (
	//"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func compareRecipe(t *testing.T, objA, objB objects.Recipe, msgAndArgs ...interface{}) (ok bool) {
	objA.CreatedAt = time.Time{}
	objB.CreatedAt = time.Time{}
	return assert.Equal(t, objA, objB, msgAndArgs)
}
func compareRecipes(t *testing.T, listA, listB []objects.Recipe, msgAndArgs ...interface{}) (ok bool) {
	for i := 0; i < len(listA); i++ {
		listA[i].CreatedAt = time.Time{}
	}
	for i := 0; i < len(listB); i++ {
		listB[i].CreatedAt = time.Time{}
	}
	return assert.ElementsMatch(t, listA, listB, msgAndArgs)
}

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
Delete grade - successful operation (recipe/account exist)
*/
func TestDelGradeRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockAcc := new(mocks.AccountsRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)
	
	objRcp := dbuilder.RecipeMother{}.Obj0()
	objAcc := dbuilder.AccountMother{}.Obj0()

	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()

	mockAcc.On("Find", objAcc.Login).Return(objAcc, nil).Once()

	mockRec.On("DeleteGrade", objRcp.Id, objAcc.Login).Return(nil).Once()

	err := allM.Recipes.DeleteGrade(objRcp.Id, objAcc.Login)

	assert.Nil(t, err, "Deletion a grade from recipe has unexpected error")
	mockRec.AssertExpectations(t)
	mockAcc.AssertExpectations(t)
}

/*
Delete recipe - successful operation (admin does action)
*/
func TestDelRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockAcc := new(mocks.AccountsRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)
	
	objRcp := dbuilder.RecipeMother{}.Obj2()
	objAdmin := dbuilder.AccountMother{}.Obj0()
	objAuthor := dbuilder.AccountMother{}.Obj1()

	mockAcc.On("Find", objAdmin.Login).Return(objAdmin, nil).Once()

	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()
	mockAcc.On("Find", objRcp.Author).Return(objAuthor, nil).Once()

	mockRec.On("Delete", objRcp.Id).Return(nil).Once()

	err := allM.Recipes.DeleteRecipe(objRcp.Id, objAdmin.Login)

	assert.Nil(t, err, "Deletion a recipe has unexpected error")
	mockRec.AssertExpectations(t)
	mockAcc.AssertExpectations(t)
}