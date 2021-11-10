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
	
	mockAcc.On("Find", obj.Author).Return(nil, nil) //FIXME: return values here
	mockRec.On("Create", obj).Return(nil).Once()

	_, err := allM.Recipes.AddRecipe(obj)
	//FIXME: 2 values: data (because auto filling, error)

	assert.Nil(t, err, "Create recipe have unexpected error")
	mockRec.AssertExpectations(t)
}