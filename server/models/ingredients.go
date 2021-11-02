package models

import (
	"api/recipes/errors"
	"api/recipes/objects"
	"api/recipes/repository"
)

type IngredientM struct {
	rep repository.IngredientsRep
	models *Models
}

func NewIngredient(rep repository.IngredientsRep, models *Models) *IngredientM {
	return &IngredientM{rep, models}
}

func (this *IngredientM) GetByRecipe(id_rcp int) ([]objects.IngredientDTO, error) {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	return this.rep.FindByRecipe(id_rcp)
}

func (this *IngredientM) PostToRecipe(id_rcp int, ing_arr *[]objects.Ingredient) error {
	/*_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}*/
	return nil
}