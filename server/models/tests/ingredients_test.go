package models_test

import (
	"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"api/recipes/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

/// CLASSIC STYLE (Stub)
/*
Check ingredient exists - ingredient exists
*/
func TestExistsIngredients(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.IngredientMother{}.All()
	objCat := dbuilder.IngredientMother{}.Obj0()

	mockRep := repository.NewIngredientsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil { panic(err) }

	model := models.NewIngredient(mockRep, nil)
	res := model.IsExists(objCat.Title)

	assert.Equal(t, res, true, "Wrong exists bool")
}
/*
ingredient not exists
*/
func TestNotExistsIngredients(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.IngredientMother{}.All()

	mockRep := repository.NewIngredientsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil { panic(err) }

	model := models.NewIngredient(mockRep, nil)
	res := model.IsExists("sheeeesh")

	assert.Equal(t, res, false, "Wrong exists bool")
}

/*
Get all ingredients - sucсess
*/
func TestGetIngredients(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {
		panic(err)
	}

	objArr := dbuilder.IngredientMother{}.All()
	objCat := dbuilder.IngredientMother{}.Obj0()

	mockRep := repository.NewIngredientsRep(db)
	err = mockRep.CreateList(objArr)
	if err != nil { panic(err) }

	model := models.NewIngredient(mockRep, nil)
	res, err := model.Get(objCat.Title)

	assert.Nil(t, err, "Get has unexpected error")
	assert.Equal(t, res, objCat, "Get has unexpected error")
}

/*
Get ingredients by recipe - ingredients and recipe with such recipe exist
*/
func TestGetByRecipeIngredients(t *testing.T) {
	db, err := stubConnecton()
	if err != nil {
		panic(err)
	}

	ingArr := dbuilder.IngredientMother{}.All()
	recArr := dbuilder.RecipeMother{}.All()
	ingRecArr := dbuilder.RecipeIngredientMother{}.All()

	mockIng := repository.NewIngredientsRep(db)
	err = mockIng.CreateList(ingArr)
	if err != nil { panic(err) }

	mockRec := repository.NewRecipesRep(db)
	err = mockRec.CreateList(recArr)
	if err != nil { panic(err) }

	for _, obj := range ingRecArr {
		mockIng.AddToRecipe(&obj)
		if err != nil {
			panic(err) 
		}
	}
	expArr := []objects.RecipeIngredient{ingRecArr[0], ingRecArr[2]}
	
	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Ingredients = models.NewIngredient(mockIng, allM)
	resArr, err := allM.Ingredients.GetByRecipe(recArr[0].Id)

	assert.Nil(t, err, "Get has unexpected error")
	assert.ElementsMatch(t, resArr, expArr)	
}



/// LONDON STYLE (Mock)

/*
Create ingredient - successful operation
*/
func TestCreateIngredient(t *testing.T) {
	mockRep := new(mocks.IngredientsRep)
	model := models.NewIngredient(mockRep, nil)
	obj := dbuilder.IngredientMother{}.Obj0()

	mockRep.On("Get", obj.Title).Return(nil, errors.RecordNotFound).Once()
	mockRep.On("Create", obj).Return(nil).Once()

	err := model.Create(obj)

	assert.Nil(t, err, "Create ingredient have unexpected error")
	mockRep.AssertExpectations(t)
}

/*
Add ingredient to recipe - successful operation, new ingredient is created
*/
func TestAddIngredientRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockIng := new(mocks.IngredientsRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Ingredients = models.NewIngredient(mockIng, allM)

	objRIng := dbuilder.RecipeIngredientMother{}.Obj0()
	objIng := &objects.Ingredient{Title: objRIng.Ingredient_id}
	objRcp := dbuilder.RecipeMother{}.Obj0()

	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()

	mockIng.On("Get", objIng.Title).Return(nil, errors.RecordNotFound).Once()

	mockIng.On("Create", objIng).Return(nil).Once()

	mockIng.On("AddToRecipe", objRIng).Return(nil).Once()

	err := allM.Ingredients.AddToRecipe(objRcp.Id, objRIng)

	assert.Nil(t, err, "Addition a new ingredient has unexpected error")
	mockRec.AssertExpectations(t)
	mockIng.AssertExpectations(t)
}


/*
Delete ingredient - successful operation (recipe/ingredient exist)
*/
func TestDelIngredientRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockIng := new(mocks.IngredientsRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Ingredients = models.NewIngredient(mockIng, allM)
	
	objRIng := dbuilder.RecipeIngredientMother{}.Obj0()
	objIng := &objects.Ingredient{Title: objRIng.Ingredient_id}
	objRcp := dbuilder.RecipeMother{}.Obj0()

	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()

	mockIng.On("Get", objIng.Title).Return(objIng, nil).Once()

	mockIng.On("DelFromRecipe", objRcp.Id, objIng.Title).Return(nil).Once()

	err := allM.Ingredients.DelFromRecipe(objRcp.Id, objRIng)

	assert.Nil(t, err, "Deletion a category from recipe has unexpected error")
	mockRec.AssertExpectations(t)
	mockIng.AssertExpectations(t)
}

/*
Post ingredient to recipe - successful operation
*/
func TestPostIngredientRecipe(t *testing.T) {
	mockRec := new(mocks.RecipesRep)
	mockIng := new(mocks.IngredientsRep)

	allM := new(models.Models)
	allM.Recipes = models.NewRecipe(mockRec, allM)
	allM.Ingredients = models.NewIngredient(mockIng, allM)

	objRIngArr := []objects.RecipeIngredient{
		*dbuilder.RecipeIngredientMother{}.Obj0(),
	}
	
	objRcp := dbuilder.RecipeMother{}.Obj0()

	mockRec.On("FindById", objRcp.Id).Return(objRcp, nil).Once()

	for _, ingredient := range(objRIngArr) {
		mockIng.On("Get", ingredient.Ingredient_id).Return(nil, errors.RecordNotFound).Once()
		mockIng.On("Create", &objects.Ingredient{Title: ingredient.Ingredient_id}).Return(nil).Once()
	}

	mockIng.On("ReplaceInRecipe", objRcp.Id, objRIngArr).Return(nil).Once()

	err := allM.Ingredients.PostToRecipe(objRcp.Id, &objRIngArr)

	assert.Nil(t, err, "Post a new ingredient/s has unexpected error")
	mockRec.AssertExpectations(t)
	mockIng.AssertExpectations(t)
}