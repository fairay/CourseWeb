package models_test

import (
	"api/recipes/errors"
	"api/recipes/mocks"
	"api/recipes/models"
	"api/recipes/objects"
	"api/recipes/objects/dbuilder"
	"testing"

	"github.com/stretchr/testify/assert"
)

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