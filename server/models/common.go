package models

import (
	"api/recipes/repository"

	"github.com/jinzhu/gorm"
)

type Models struct {
	Recipes *RecipeM
	Accounts *AccountM
	Category *CategoryM
	Ingredients *IngredientM
	Steps *StepM
}

func InitModels(db *gorm.DB) *Models {
	models := new(Models)

	models.Recipes = NewRecipe(repository.NewRecipesRep(db), models)
	models.Accounts = NewAccount(repository.NewAccountsRep(db), models)
	models.Category = NewCategory(repository.NewCategotiesRep(db), models)
	models.Ingredients = NewIngredient(repository.NewIngredientsRep(db), models)
	models.Steps = NewStep(repository.NewStepsRep(db), models)

	return models
}