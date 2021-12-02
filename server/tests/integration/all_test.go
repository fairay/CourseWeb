package integration_test

import (
	"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"api/recipes/repository"
	"api/recipes/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

/// Integration with only two models tested (without mocking/stubing others)
/* Get author
- account exists
*/
func TestGetAuthor(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	objRcp := dbuilder.RecipeMother{}.Obj0()
	objAcc := dbuilder.AccountMother{}.Obj0()
	resAuthor := new(objects.Account)

	mockRep := repository.NewRecipesRep(db)
	mockAcc := repository.NewAccountsRep(db)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	var errArr [3]error
	errArr[0] = allM.Accounts.Create(objAcc)
	errArr[1] = allM.Recipes.AddRecipe(objRcp)
	resAuthor, errArr[2] = allM.Recipes.GetAuthor(objRcp.Id)

	for i := 0; i < len(errArr); i++ {
		assert.Nil(t, errArr[i], "Unexpected error")
	}
	assert.Equal(t, resAuthor, objAcc, "GetAuthor has unexpected error")
}

/* Get author
- account not exists
*/

/// Integration with only two models tested (with mocking/stubing others)
/* Get categories by recipe
- category and recipe with such recipe exist
*/
func TestGetByRecipeCategory(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	accArr := dbuilder.AccountMother{}.All()
	catArr := dbuilder.CategoryMother{}.All()
	recArr := dbuilder.RecipeMother{}.All()

	mockAcc := new(mocks.AccountsRep)
	mockCat := repository.NewCategotiesRep(db)
	mockRec := repository.NewRecipesRep(db)
	mockAcc.On("Find", recArr[0].Author).Return(&accArr[0], nil).Twice()
	expArr := []objects.Category{catArr[0]}

	allM := new(models.Models)
	allM.Accounts = models.NewAccount(mockAcc, allM)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Category = models.NewCategory(mockCat, allM)

	var errArr [6]error
	var resArr []objects.Category
	errArr[0] = allM.Recipes.AddRecipe(&recArr[0])
	errArr[1] = allM.Recipes.AddRecipe(&recArr[1])

	errArr[2] = allM.Category.AddToRecipe(recArr[0].Id, &catArr[0])
	errArr[3] = allM.Category.AddToRecipe(recArr[1].Id, &catArr[0])
	errArr[4] = allM.Category.AddToRecipe(recArr[1].Id, &catArr[1])

	resArr, errArr[5] = allM.Category.GetByRecipe(recArr[0].Id)

	for i := 0; i < len(errArr); i++ {
		assert.Nil(t, errArr[i], "Unexpected error")
	}
	assert.ElementsMatch(t, resArr, expArr)
	mockAcc.AssertExpectations(t)
}

/* Get categories by recipe
- such recipe doesn't exist
*/
func TestGetByRecipeCategory_WrongRecipe(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {
		panic(err)
	}

	accArr := dbuilder.AccountMother{}.All()
	catArr := dbuilder.CategoryMother{}.All()
	recArr := dbuilder.RecipeMother{}.All()

	mockAcc := new(mocks.AccountsRep)
	mockCat := repository.NewCategotiesRep(db)
	mockRec := repository.NewRecipesRep(db)
	mockAcc.On("Find", recArr[0].Author).Return(&accArr[0], nil).Twice()
	expArr := []objects.Category{}

	allM := new(models.Models)
	allM.Accounts = models.NewAccount(mockAcc, allM)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Category = models.NewCategory(mockCat, allM)

	var errArr [6]error; expErrArr := [6]error{nil, nil, nil, nil, nil, errors.UnknownRecipe}
	var resArr []objects.Category
	errArr[0] = allM.Recipes.AddRecipe(&recArr[0])
	errArr[1] = allM.Recipes.AddRecipe(&recArr[1])

	errArr[2] = allM.Category.AddToRecipe(recArr[0].Id, &catArr[0])
	errArr[3] = allM.Category.AddToRecipe(recArr[1].Id, &catArr[0])
	errArr[4] = allM.Category.AddToRecipe(recArr[1].Id, &catArr[1])

	resArr, errArr[5] = allM.Category.GetByRecipe(-1)

	for i := 0; i < len(errArr); i++ {
		assert.Equal(t, expErrArr[i], errArr[i], "Unexpected error")
	}
	assert.ElementsMatch(t, resArr, expArr)
	mockAcc.AssertExpectations(t)
}