package models_test

import (
	//"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects/dbuilder"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Create step - successful operation
*/
func TestCreateStep(t *testing.T) {
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
	
	fmt.Println(objAcc, objRec, obj)

	mockAcc.On("Find", objAcc.Login).Return(nil, nil) //FIXME: return values here

	mockRec.On("FindById", objRec.Id).Return(nil, nil)  //FIXME: return values here
	mockAcc.On("Find", objAcc.Login).Return(nil, nil) //TODO: переписать код функций, второй раз вызывается

	mockRec.On("FindById", objRec.Id).Return(nil, nil)  //TODO: встречается второй раз, но без этого кажется никак

	mockRec.On("Create", obj).Return(nil).Once()

	mockRec.On("FindById", objRec.Id).Return(nil, nil) //TODO: третий раз, скорее всего можно убрать эту проверку из функции
	mockStp.On("FindStepLast", obj.Recipe).Return(nil, nil)

	_, err := allM.Steps.AddStep(objRec.Id, obj, objAcc.Login)
	//FIXME: 2 values: data (because auto filling, error)

	assert.Nil(t, err, "Create step have unexpected error")
	mockRec.AssertExpectations(t)
}