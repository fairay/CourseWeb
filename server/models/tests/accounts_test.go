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
Create account - successful operation
 */
func TestCreateAccount(t *testing.T) {
	mockRep := new(mocks.AccountsRep)
	model := models.NewAccount(mockRep, nil)
	obj := dbuilder.AccountMother{}.Obj0()

	mockRep.On("Find", obj.Login).Return(nil, errors.RecordNotFound).Once()
	mockRep.On("Create", obj).Return(nil).Once()

	err := model.Create(obj)

	assert.Nil(t, err, "Create account have unexpected error")
	mockRep.AssertExpectations(t)
}

/*
Update role by admin- successful operation
*/
func TestUpdateRole(t *testing.T) {
	mockRep := new(mocks.AccountsRep)
	model := models.NewAccount(mockRep, nil)
	objCur := dbuilder.AccountMother{}.Obj0()
	objGoal := dbuilder.AccountMother{}.Obj1()
	objUdp := dbuilder.AccountMother{}.ObjUdp()

	mockRep.On("Find", objCur.Login).Return(objCur, nil).Once()
	mockRep.On("Find", objGoal.Login).Return(objGoal, nil).Once()
	mockRep.On("UpdateRole", objGoal.Login, objUdp.Role).Return(nil).Once()

	err := model.UpdateRole(objCur.Login, objGoal.Login, objUdp.Role)

	assert.Nil(t, err, "Update account have unexpected error")
	mockRep.AssertExpectations(t)
}