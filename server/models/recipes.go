package models

import (
	"api/recipes/objects"
	"api/recipes/repository"
)

type RecipeM struct {
	rep repository.RecipesRep
	models *Models
}

func NewRecipe(rep repository.RecipesRep, models *Models) *RecipeM {
	return &RecipeM{rep, models}
}

func (this *RecipeM) GetAll() ([]objects.Recipe) {
	temp := this.rep.List()
	return temp
}

func (this *RecipeM) AddRecipe(obj *objects.Recipe) (err error) {
	_, err = this.models.Accounts.find(obj.Author);
	if err != nil {
		return err
	}

	/* TODO: check for right categories (возможно и не надо, если из
	выпадающего списка...)

	TODO: validate other data
	*/
	err = this.rep.Create(obj)
	return
}