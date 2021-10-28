package models

import (
	"api/recipes/objects"
	"api/recipes/repository"
)

type RecipeM struct {
	rep repository.RecipesRep
}

func NewRecipe(rep repository.RecipesRep) *RecipeM {
	return &RecipeM{rep}
}

func (this *RecipeM) GetAll() ([]objects.Recipe) {
	temp := this.rep.List()
	return temp
}