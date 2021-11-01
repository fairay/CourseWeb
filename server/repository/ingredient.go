package repository

import (
	"api/recipes/objects"
	//"strings"

	"github.com/jinzhu/gorm"
)

type IngredientsRep interface {
	FindByRecipe(id_rcp int) ([]objects.IngredientDTO, error)
}

type PGIngredientsRep struct {
	db *gorm.DB
}

func NewIngredientsRep(db *gorm.DB) *PGIngredientsRep {
	return &PGIngredientsRep{db}
}

func (this *PGIngredientsRep) FindByRecipe(id_rcp int) ([]objects.IngredientDTO, error) {
	temp := []objects.RecipeIngredient{}
	//recipe := objects.Recipe{Id: id_rcp}

	//err := this.db.Model(&recipe).Association("Ingredients").Find(&temp).Error
	err := this.db.Preload("Ingredients").Preload("Recipes").Find(&temp).Error

	return nil, err
}