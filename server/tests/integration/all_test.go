package integration_test

import (
	"api/recipes/models"
	"api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"api/recipes/repository"
	"api/recipes/tests"
	"api/recipes/errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

/// Integration with only two models tested (without mocking/stubing others)
/* Get author
- account exists
*/
func TestGetAuthor(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil {	panic(err) }

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
func TestGetAuthor_NotRec(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil { panic(err) }

	objRcp := dbuilder.RecipeMother{}.Obj0()
	resAuthor := new(objects.Account)
	var objAuthor *objects.Account = nil

	mockRep := repository.NewRecipesRep(db)
	mockAcc := repository.NewAccountsRep(db)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	var errArr [2]error
	errArr[0] = allM.Recipes.AddRecipe(objRcp)
	resAuthor, errArr[1] = allM.Recipes.GetAuthor(objRcp.Id)

	assert.Equal(t, errArr[0], errors.UnknownAccount)
	assert.Equal(t, errArr[1], errors.UnknownRecipe)
	assert.Equal(t, objAuthor, resAuthor, "GetAuthor has unexpected error")
}


/* Get recipe
- recipe not exists
*/
func TestGetRecipe_NotRec(t *testing.T) {
	db, err := tests.StubConnecton()
	if err != nil { panic(err) }

	objRcp := dbuilder.RecipeMother{}.Obj0()
	objAcc := dbuilder.AccountMother{}.Obj0()
	resRecipe := new(objects.Recipe)
	var resRcp *objects.Recipe = nil

	mockRep := repository.NewRecipesRep(db)
	mockAcc := repository.NewAccountsRep(db)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRep, allM)
	allM.Accounts = models.NewAccount(mockAcc, allM)

	var errArr [2]error
	errArr[0] = allM.Accounts.Create(objAcc)
	resRecipe, errArr[1] = allM.Recipes.FindById(objRcp.Id)

	assert.Nil(t, errArr[0], "Create account has unexpected error")
	assert.Equal(t, errArr[1], errors.UnknownRecipe)
	assert.Equal(t, resRcp, resRecipe, "GetAuthor has unexpected error")
}


/// Integration with only two models tested (with mocking/stubing others)
