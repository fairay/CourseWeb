package models_test

import (
	"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects/dbuilder"
	"api/recipes/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

/// CLASSIC STYLE (Stub)

/*
Get account - account exists
*/
func TestFindAccount(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.AccountMother{}.All()
	objAcc := dbuilder.AccountMother{}.Obj0()

	mockRep := repository.NewAccountsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil { panic(err) }

	model := models.NewAccount(mockRep, nil)
	res, err := model.Find(objAcc.Login)

	assert.Nil(t, err, "Find has unexpected error")
	assert.Equal(t, res, objAcc, "Find has unexpected error")
}

/*
Get role - account exists
*/
func TestGetRoleAccount(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.AccountMother{}.All()
	objAcc := dbuilder.AccountMother{}.Obj0()

	mockRep := repository.NewAccountsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil { panic(err) }

	model := models.NewAccount(mockRep, nil)
	resRole, err := model.GetRole(objAcc.Login)

	assert.Nil(t, err, "GetRole has unexpected error")
	assert.Equal(t, resRole, objAcc.Role, "GetRole has unexpected error")
}

/*
Account exists
*/
func TestLogIn(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.AccountMother{}.All()
	objAcc := dbuilder.AccountMother{}.Obj0()

	mockRep := repository.NewAccountsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil { panic(err) }

	model := models.NewAccount(mockRep, nil)
	res, err := model.LogIn(objAcc.Login, objAcc.HashedPassword)

	assert.Nil(t, err, "GetRole has unexpected error")
	assert.Equal(t, res, objAcc, "GetRole has unexpected error")
}

/// LONDON STYLE (Mock)

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
