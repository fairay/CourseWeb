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

//TODO: return array and error? or only array
func (this *RecipeM) FindByLogin(login string) ([]objects.Recipe) {
	isExist := this.models.Accounts.IsExists(login)
	if isExist == false {
		return nil
	}

	temp := this.rep.FindByLogin(login)
	return temp
}

func (this *RecipeM) AddRecipe(obj *objects.Recipe) (err error) {
	_, err = this.models.Accounts.find(obj.Author);
	if err != nil {
		return err
	}

	// TODO: validate other data
	err = this.rep.Create(obj)
	return
}