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
Create step (by recipe admin) - successful operation
*/
func TestCreateStepAdmin(t *testing.T) {
	mockStp := new(mocks.StepsRep)
	mockRec := new(mocks.RecipesRep)
	mockAcc := new(mocks.AccountsRep)

	allM := new(models.Models)
	allM.Steps = models.NewStep(mockStp, allM)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	objAcc := dbuilder.AccountMother{}.Obj0()
	objRec := dbuilder.RecipeMother{}.Obj0()
	obj := dbuilder.StepMother{}.Obj0()
	
	mockAcc.On("Find", objAcc.Login).Return(objAcc, nil).Once() 
	mockRec.On("FindById", objRec.Id).Return(objRec, nil).Once()
	mockStp.On("Create", obj).Return(nil).Once()

	err := allM.Steps.AddStep(objRec.Id, obj, objAcc.Login)

	assert.Nil(t, err, "Create step have unexpected error")
	mockRec.AssertExpectations(t)
}

/*
Create step (by recipe author) - successful operation
*/
func TestCreateStepAuthor(t *testing.T) {
	mockStp := new(mocks.StepsRep)
	mockRec := new(mocks.RecipesRep)
	mockAcc := new(mocks.AccountsRep)

	allM := new(models.Models)
	allM.Steps = models.NewStep(mockStp, allM)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	objAcc := dbuilder.AccountMother{}.Obj1()
	objRec := dbuilder.RecipeMother{}.Obj2()
	obj := dbuilder.StepMother{}.Obj2()
	
	mockAcc.On("Find", objAcc.Login).Return(objAcc, nil).Twice() 
	mockRec.On("FindById", objRec.Id).Return(objRec, nil).Twice()
	mockStp.On("Create", obj).Return(nil).Once()

	err := allM.Steps.AddStep(objRec.Id, obj, objAcc.Login)

	assert.Nil(t, err, "Create step have unexpected error")
	mockRec.AssertExpectations(t)
}
