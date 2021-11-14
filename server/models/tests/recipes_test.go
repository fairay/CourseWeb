package models_test

import (
	//"api/recipes/errors"
	"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"api/recipes/repository"
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

/// CLASSIC STYLE (Stub)

/* Get all recipes
- sucсess
*/
func TestGetAllRecipes(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	objArr := dbuilder.RecipeMother{}.All()

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil {	panic(err) }

	model := models.NewRecipe(mockRep, nil)

	resArr := model.GetAll()

	compareRecipes(t, resArr, objArr)
}
/* Get all recipes
- empty recipes
*/
func TestGetAllRecipes_Empty(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	objArr := []objects.Recipe{}

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil { panic(err) }

	model := models.NewRecipe(mockRep, nil)

	resArr := model.GetAll()

	compareRecipes(t, resArr, objArr)
}

/* Get author
- account exists
*/
func TestGetAuthor(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	objAccArr := dbuilder.AccountMother{}.All()
	objRcp := dbuilder.RecipeMother{}.Obj0()
	objAcc := &objAccArr[0]

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil { panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	err = mockAcc.CreateList(objAccArr)
	if err != nil { panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	resAuthor, err := allM.Recipes.GetAuthor(objRcp.Id)

	assert.Nil(t, err, "GetAuthor has unexpected error")
	assert.Equal(t, resAuthor, objAcc, "GetAuthor has unexpected error")
}
/* Get author
- recipe not exists
*/
func TestGetAuthor_NotAcc(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	objAccArr := dbuilder.AccountMother{}.All()
	// objRcp := dbuilder.RecipeMother{}.Obj0()
	var objAcc *objects.Account = nil;

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil { panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	err = mockAcc.CreateList(objAccArr)
	if err != nil { panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	resAuthor, err := allM.Recipes.GetAuthor(-1)

	assert.Equal(t, errors.UnknownRecipe, err, "GetAuthor have unexpected error")
	assert.Equal(t, resAuthor, objAcc, "GetAuthor has unexpected error")
}
/* Get author
- account not exists
*/
func TestGetAuthor_NotRec(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	// objAccArr := dbuilder.AccountMother{}.All()
	objRcp := dbuilder.RecipeMother{}.Obj0()
	var objAcc *objects.Account = nil;

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil { panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	// err = mockAcc.CreateList(objAccArr)
	// if err != nil { panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	resAuthor, err := allM.Recipes.GetAuthor(objRcp.Id)

	assert.Equal(t, errors.UnknownAccount, err, "GetAuthor have unexpected error")
	assert.Equal(t, resAuthor, objAcc, "GetAuthor has unexpected error")
}


/* Find recipes by login
- sucсess
*/
func TestFindRecipesByLogin(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	objAccArr := dbuilder.AccountMother{}.All()
	objAcc := dbuilder.AccountMother{}.Obj0()
	aimArr := []objects.Recipe{objRcpArr[0], objRcpArr[1]}

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil { panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	err = mockAcc.CreateList(objAccArr)
	if err != nil { panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	resArr, _ := allM.Recipes.FindByLogin(objAcc.Login)

	compareRecipes(t, resArr, aimArr)
}
/* Find recipes by login
- sucсess with empty table
*/
func TestFindRecipesByLogin_Empty(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	// objRcpArr := dbuilder.RecipeMother{}.All()
	objAccArr := dbuilder.AccountMother{}.All()
	objAcc := dbuilder.AccountMother{}.Obj0()
	aimArr := []objects.Recipe{}

	mockRep := repository.NewRecipesRep(db)
	// err = mockRep.CreateList(objRcpArr)
	// if err != nil { panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	err = mockAcc.CreateList(objAccArr)
	if err != nil { panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	resArr, _ := allM.Recipes.FindByLogin(objAcc.Login)

	compareRecipes(t, resArr, aimArr)
}
/* Find recipes by login
- account not exists
*/
func TestFindRecipesByLogin_NotExists(t *testing.T) {
	db, err := stubConnecton()
	if err != nil { panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	// objAccArr := dbuilder.AccountMother{}.All()
	objAcc := dbuilder.AccountMother{}.Obj0()
	aimArr := []objects.Recipe{}

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil { panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	// err = mockAcc.CreateList(objAccArr)
	// if err != nil { panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	resArr, err := allM.Recipes.FindByLogin(objAcc.Login)

	assert.Equal(t, errors.UnknownAccount, err, "Find recipes by login have unexpected error")
	compareRecipes(t, resArr, aimArr)
}


/* Find recipes by its id
- sucсess
*/
func TestFindRecipesById(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	aimRcp := dbuilder.RecipeMother{}.Obj0()

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil {	panic(err) }

	model := models.NewRecipe(mockRep, nil)

	resRcp, err := model.FindById(aimRcp.Id)

	assert.Nil(t, err, "Find recipes by its id has unexpected error")
	compareRecipe(t, *aimRcp, *resRcp)
}
/* Find recipes by its id
- not found
*/
func TestFindRecipesById_Empty(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	// aimRcp := dbuilder.RecipeMother{}.Obj0()

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil {	panic(err) }

	model := models.NewRecipe(mockRep, nil)

	res, err := model.FindById(-1)

	var nil_ptr *objects.Recipe = nil;
	assert.Equal(t, errors.UnknownRecipe, err, "Find recipes by its id has unexpected error")
	assert.Equal(t, nil_ptr, res, "Find recipes by its id has unexpected result")
}

/* Get amount of grades
- 2 likes
*/
func TestGetAmountGrades_2(t *testing.T) {
	db, err := stubConnecton()
	if err != nil { panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	objAccArr := dbuilder.AccountMother{}.All()
	objRcp := dbuilder.RecipeMother{}.Obj0()
	objAcc1 := dbuilder.AccountMother{}.Obj1()
	objAcc2 := dbuilder.AccountMother{}.Obj2()
	aimNum := 2

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil { panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	err = mockAcc.CreateList(objAccArr)
	if err != nil { panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	_ = allM.Recipes.AddGrade(objRcp.Id, objAcc1.Login)
	_ = allM.Recipes.AddGrade(objRcp.Id, objAcc2.Login)

	resNum, err := allM.Recipes.GetAmountGrades(objRcp.Id)

	assert.Nil(t, err, "GetAmountGrades has unexpected error")
	assert.Equal(t, aimNum, resNum, "GetAmountGrades returns wrong answer")
}
/* Get amount of grades
- 0 likes
*/
func TestGetAmountGrades_0(t *testing.T) {
	db, err := stubConnecton()
	if err != nil { panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	objAccArr := dbuilder.AccountMother{}.All()
	objRcp := dbuilder.RecipeMother{}.Obj0()
	aimNum := 0

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil { panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	err = mockAcc.CreateList(objAccArr)
	if err != nil { panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)
	resNum, err := allM.Recipes.GetAmountGrades(objRcp.Id)

	assert.Nil(t, err, "GetAmountGrades has unexpected error")
	assert.Equal(t, aimNum, resNum, "GetAmountGrades returns wrong answer")
}
/* Get amount of grades
- recipe doesn't exist
*/
func TestGetAmountGrades_None(t *testing.T) {
	db, err := stubConnecton()
	if err != nil { panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	objAccArr := dbuilder.AccountMother{}.All()
	aimNum := 0

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil { panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	err = mockAcc.CreateList(objAccArr)
	if err != nil { panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)
	resNum, err := allM.Recipes.GetAmountGrades(-1)

	assert.Equal(t, errors.UnknownRecipe, err, "GetAmountGrades has unexpected error")
	assert.Equal(t, aimNum, resNum, "GetAmountGrades returns wrong answer")
}


/* Get liked by login
- 1 like, account/recipe exist
*/
func TestGetLikedByLogin_1(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	objAccArr := dbuilder.AccountMother{}.All()
	objAcc := dbuilder.AccountMother{}.Obj1()

	aimArr := []objects.Recipe{objRcpArr[0]}

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil {	panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	err = mockAcc.CreateList(objAccArr)
	if err != nil {	panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	for i := 0; i < len(aimArr); i++ {
		_ = allM.Recipes.AddGrade(aimArr[i].Id, objAcc.Login)
	}

	resArr, err := allM.Recipes.GetLikedByLogin(objAcc.Login)

	compareRecipes(t, resArr, aimArr)
}
/* Get liked by login
- 0 like, account/recipe exist
*/
func TestGetLikedByLogin_0(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	objAccArr := dbuilder.AccountMother{}.All()
	objAcc := dbuilder.AccountMother{}.Obj1()

	aimArr := []objects.Recipe{}

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil {	panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	err = mockAcc.CreateList(objAccArr)
	if err != nil {	panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	for i := 0; i < len(aimArr); i++ {
		_ = allM.Recipes.AddGrade(aimArr[i].Id, objAcc.Login)
	}

	resArr, err := allM.Recipes.GetLikedByLogin(objAcc.Login)

	compareRecipes(t, resArr, aimArr)
}
/* Get liked by login
- 0 like, account/recipe exist
*/
func TestGetLikedByLogin_None(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {	panic(err) }

	objRcpArr := dbuilder.RecipeMother{}.All()
	objAccArr := dbuilder.AccountMother{}.All()
	objAcc := dbuilder.AccountMother{}.Obj1()

	aimArr := []objects.Recipe{}

	mockRep := repository.NewRecipesRep(db)
	err = mockRep.CreateList(objRcpArr)
	if err != nil {	panic(err) }

	mockAcc := repository.NewAccountsRep(db)
	err = mockAcc.CreateList(objAccArr)
	if err != nil {	panic(err) }

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	for i := 0; i < len(aimArr); i++ {
		_ = allM.Recipes.AddGrade(aimArr[i].Id, objAcc.Login)
	}

	resArr, err := allM.Recipes.GetLikedByLogin(nWord)

	assert.Equal(t, errors.UnknownAccount, err, "GetLikedByLogin has unexpected error")
	compareRecipes(t, resArr, aimArr)
}


/// LONDON STYLE (Mock)

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

	assert.Nil(t, err, "Create recipe has unexpected error")
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
