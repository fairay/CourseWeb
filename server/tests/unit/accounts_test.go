package models_test

import (
	"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"api/recipes/repository"
	"api/recipes/tests"
	err "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

/// CLASSIC STYLE (Stub)

/* Find account
- account exists
*/
func TestFindAccount(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.AccountMother{}.All()
	objAcc := &objArr[1]

	mockRep := repository.NewAccountsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewAccount(mockRep, nil)
	res, err := model.Find(objAcc.Login)

	assert.Nil(t, err, "Find has unexpected error")
	assert.Equal(t, objAcc, res, "Find has unexpected result")
}

/* Find account
- account doesn't exists
*/
func TestFindNoneAccount(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.AccountMother{}.All()

	mockRep := repository.NewAccountsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewAccount(mockRep, nil)
	res, err := model.Find(nWord)

	var nil_ptr *objects.Account = nil
	assert.Equal(t, errors.RecordNotFound, err, "Find has unexpected error")
	assert.Equal(t, nil_ptr, res, "Find has unexpected result")
}

/* Get role
- account exists
*/
func TestGetRoleAccount(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.AccountMother{}.All()
	objAcc := dbuilder.AccountMother{}.Obj0()

	mockRep := repository.NewAccountsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewAccount(mockRep, nil)
	resRole, err := model.GetRole(objAcc.Login)

	assert.Nil(t, err, "GetRole has unexpected error")
	assert.Equal(t, resRole, objAcc.Role, "GetRole has unexpected error")
}

/* Get role
- account doesn't exists
*/
func TestGetRoleNoneAccount(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.AccountMother{}.All()

	mockRep := repository.NewAccountsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewAccount(mockRep, nil)
	resRole, err := model.GetRole(nWord)

	assert.Equal(t, errors.RecordNotFound, err, "GetRole has unexpected error")
	assert.Equal(t, "", resRole, "GetRole has unexpected result")
}

/* Log in
- account exists, correct password
*/
func TestLogIn(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.AccountMother{}.All()
	objAcc := &objArr[0] // not dbuilder.AccountMother{}.Obj0(), hash generated differently every time

	mockRep := repository.NewAccountsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewAccount(mockRep, nil)
	res, err := model.LogIn(objAcc.Login, objAcc.Login)

	assert.Nil(t, err, "Login has unexpected error")
	assert.Equal(t, res, objAcc, "Login has unexpected result")
}

/* Log in
- account exists, wrong password
*/
func TestLogInWrongPW(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.AccountMother{}.All()
	objAcc := &objArr[0]

	mockRep := repository.NewAccountsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewAccount(mockRep, nil)
	res, err := model.LogIn(objAcc.Login, nWord)

	var nil_ptr *objects.Account = nil
	assert.Equal(t, errors.WrongPassword, err, "Login has unexpected error")
	assert.Equal(t, nil_ptr, res, "Login has unexpected result")
}

/* Log in
- account exists, wrong password
*/
func TestLogInNone(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.AccountMother{}.All()

	mockRep := repository.NewAccountsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {
		panic(err)
	}

	model := models.NewAccount(mockRep, nil)
	res, err := model.LogIn(nWord, nWord)

	var nil_ptr *objects.Account = nil
	assert.Equal(t, errors.RecordNotFound, err, "Login has unexpected error")
	assert.Equal(t, nil_ptr, res, "Login has unexpected result")
}

/// LONDON STYLE (Mock)

/* Create account
- successful operation
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

/* Create account
- already existing account
*/
func TestCreateExistsAccount(t *testing.T) {
	mockRep := new(mocks.AccountsRep)
	model := models.NewAccount(mockRep, nil)
	obj := dbuilder.AccountMother{}.Obj0()

	mockRep.On("Find", obj.Login).Return(obj, nil).Once()
	// mockRep.On("Create", obj).Return().Once()

	err := model.Create(obj)

	assert.Equal(t, errors.AccountExists, err, "Create account have unexpected error")
	mockRep.AssertExpectations(t)
}

/* Create account
- creation failed
*/
func TestCreateFailAccount(t *testing.T) {
	mockRep := new(mocks.AccountsRep)
	model := models.NewAccount(mockRep, nil)
	obj := dbuilder.AccountMother{}.Obj0()

	mockRep.On("Find", obj.Login).Return(nil, errors.RecordNotFound).Once()
	mockRep.On("Create", obj).Return(err.New("")).Once()

	err := model.Create(obj)

	assert.Equal(t, errors.DBAdditionError, err, "Create account have unexpected error")
	mockRep.AssertExpectations(t)
}

/* Update role by admin
- successful operation
*/
func TestUpdateRole(t *testing.T) {
	mockRep := new(mocks.AccountsRep)
	model := models.NewAccount(mockRep, nil)
	objCur := dbuilder.AccountMother{}.Obj0()
	objGoal := dbuilder.AccountMother{}.Obj1()
	objUdp := dbuilder.AccountMother{}.Obj1Udp()

	mockRep.On("Find", objCur.Login).Return(objCur, nil).Once()
	mockRep.On("Find", objGoal.Login).Return(objGoal, nil).Once()
	mockRep.On("UpdateRole", objGoal.Login, objUdp.Role).Return(nil).Once()

	err := model.UpdateRole(objCur.Login, objGoal.Login, objUdp.Role)

	assert.Nil(t, err, "Update account have unexpected error")
	mockRep.AssertExpectations(t)
}

/* Update role by admin
- admin account wasn't found
*/
func TestUpdateRole_ActorMissing(t *testing.T) {
	mockRep := new(mocks.AccountsRep)
	model := models.NewAccount(mockRep, nil)
	objCur := dbuilder.AccountMother{}.Obj0()
	objGoal := dbuilder.AccountMother{}.Obj1()
	objUdp := dbuilder.AccountMother{}.Obj1Udp()

	mockRep.On("Find", objCur.Login).Return(nil, errors.RecordNotFound).Once()
	// mockRep.On("Find", objGoal.Login).Return(objGoal, nil).Once()
	// mockRep.On("UpdateRole", objGoal.Login, objUdp.Role).Return(nil).Once()

	err := model.UpdateRole(objCur.Login, objGoal.Login, objUdp.Role)

	assert.Equal(t, errors.UnknownAccount, err, "Update account have unexpected error")
	mockRep.AssertExpectations(t)
}

/* Update role by admin
- actor is not admin
*/
func TestUpdateRole_NotAdmin(t *testing.T) {
	mockRep := new(mocks.AccountsRep)
	model := models.NewAccount(mockRep, nil)
	objCur := dbuilder.AccountMother{}.Obj1()
	objGoal := dbuilder.AccountMother{}.Obj0()
	objUdp := dbuilder.AccountMother{}.Obj1Udp()

	mockRep.On("Find", objCur.Login).Return(objCur, nil).Once()
	// mockRep.On("Find", objGoal.Login).Return(objGoal, nil).Once()
	// mockRep.On("UpdateRole", objGoal.Login, objUdp.Role).Return(nil).Once()

	err := model.UpdateRole(objCur.Login, objGoal.Login, objUdp.Role)

	assert.Equal(t, errors.AccessDenied, err, "Update account have unexpected error")
	mockRep.AssertExpectations(t)
}

/* Update role by admin
- updating account not found
*/
func TestUpdateRole_NotFound(t *testing.T) {
	mockRep := new(mocks.AccountsRep)
	model := models.NewAccount(mockRep, nil)
	objCur := dbuilder.AccountMother{}.Obj0()
	objGoal := dbuilder.AccountMother{}.Obj1()
	objUdp := dbuilder.AccountMother{}.Obj1Udp()

	mockRep.On("Find", objCur.Login).Return(objCur, nil).Once()
	mockRep.On("Find", objGoal.Login).Return(nil, errors.RecordNotFound).Once()
	// mockRep.On("UpdateRole", objGoal.Login, objUdp.Role).Return(nil).Once()

	err := model.UpdateRole(objCur.Login, objGoal.Login, objUdp.Role)

	assert.Equal(t, errors.UnknownAccount, err, "Update account have unexpected error")
	mockRep.AssertExpectations(t)
}
