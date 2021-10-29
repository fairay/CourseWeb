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

func (this *RecipeM) AddRecipe(login string) error{
	acc, err := this.models.accounts.rep.Find(login)
	
	if err != nil {
		return err
	}

	/* TODO: check for right categories (возможно и не надо, если из
	/* выпадающего списка...)
	*/



	//err = repository.RecipesRep.Create()

}