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

/*
Delete step - successful operation (admin does action)
*/
func TestDelStep(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockAcc := new(mocks.AccountsRep)
	mockStp := new(mocks.StepsRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)
	allM.Steps = models.NewStep(mockStp, allM)
	
	objRcp := dbuilder.RecipeMother{}.Obj2()
	objAdmin := dbuilder.AccountMother{}.Obj0()
	objAuthor := dbuilder.AccountMother{}.Obj1()
	objStep := dbuilder.StepMother{}.Obj2()

	mockAcc.On("Find", objAdmin.Login).Return(objAdmin, nil).Once()

	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()
	mockAcc.On("Find", objRcp.Author).Return(objAuthor, nil).Once()

	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()
	mockStp.On("FindStepByNum", objRcp.Id, objStep.Num).Return(*objStep, nil).Once()

	mockStp.On("Delete", objRcp.Id, objStep.Num).Return(nil).Once()

	err := allM.Steps.DeleteStep(objRcp.Id, objStep.Num, objAdmin.Login)

	assert.Nil(t, err, "Deletion a step has unexpected error")
	mockRec.AssertExpectations(t)
	mockAcc.AssertExpectations(t)
	mockStp.AssertExpectations(t)
}

/*
Update step - successful operation (admin does action)
*/
func TestUpdateStep(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockAcc := new(mocks.AccountsRep)
	mockStp := new(mocks.StepsRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)
	allM.Steps = models.NewStep(mockStp, allM)
	
	objRcp := dbuilder.RecipeMother{}.Obj2()
	objAdmin := dbuilder.AccountMother{}.Obj0()
	objAuthor := dbuilder.AccountMother{}.Obj1()
	objStep := dbuilder.StepMother{}.Obj2()

	mockAcc.On("Find", objAdmin.Login).Return(objAdmin, nil).Twice()

	mockAcc.On("Find", objRcp.Author).Return(objAuthor, nil).Once()
	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()
	
	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()
	mockStp.On("FindStepByNum", objRcp.Id, objStep.Num).Return(*objStep, nil).Once()

	mockStp.On("UpdateStep", objRcp.Id, objStep.Num, objStep).Return(nil).Once()

	err := allM.Steps.UpdateStep(objAdmin.Login, objRcp.Id, objStep.Num, objStep)

	assert.Nil(t, err, "Update step has unexpected error")
	mockRec.AssertExpectations(t)
	mockAcc.AssertExpectations(t)
	mockStp.AssertExpectations(t)
}