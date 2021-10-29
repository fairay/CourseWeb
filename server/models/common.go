package models

import (
	"api/recipes/repository"

	"github.com/jinzhu/gorm"
)

type Models struct {
	recipes *RecipeM
	accounts *AccountM
}

func InitModels(db *gorm.DB) *Models {
	models := &Models{}

	models.recipes = NewRecipe(repository.NewRecipesRep(db), models)
	models.accounts = NewAccount(repository.NewAccountsRep(db))

	return models
}