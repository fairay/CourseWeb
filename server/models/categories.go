package models

import (
	"api/recipes/objects"
	"api/recipes/repository"
	"api/recipes/errors"
)

type CategoryM struct {
	rep repository.CategoriesRep
	models *Models
}

func NewCategory(rep repository.CategoriesRep, models *Models) *CategoryM {
	return &CategoryM{rep, models}
}

func (this *CategoryM) GetAll() ([]objects.Category) {
	temp := this.rep.List()
	return temp
}

func (this *CategoryM) Find(ctg string) ([]objects.Category, error) {
	return this.rep.Find(ctg)
}

func (this *CategoryM) Get(ctg string) (objects.Category) {
	temp := this.rep.Get(ctg)
	return temp
}

func (this *CategoryM) GetRecipes(ctg string) ([]objects.Recipe, error) {
	if this.IsExists(ctg) == false {
		return nil, errors.UnknownRecipe
	}

	return this.rep.FindRecipes(ctg)
}

func (this *CategoryM) GetByRecipe(id_rcp int) ([]objects.Category, error) {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	return this.rep.FindByRecipe(id_rcp)
}

func (this *CategoryM) IsExists(ctg string) bool {
	data, _ := this.Find(ctg)

	if len(data) == 0 { return false }

	return true
}
