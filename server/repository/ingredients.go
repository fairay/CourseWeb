package repository

import (
	"api/recipes/objects"
	//"strings"

	"github.com/jinzhu/gorm"
)

type IngredientsRep interface {
	FindByRecipe(id_rcp int) ([]objects.RecipeIngredient, error)
}

type PGIngredientsRep struct {
	db *gorm.DB
}

func NewIngredientsRep(db *gorm.DB) *PGIngredientsRep {
	return &PGIngredientsRep{db}
}

func (this *PGIngredientsRep) FindByRecipe(idRecipe int) ([]objects.RecipeIngredient, error) {
	temp := []objects.RecipeIngredient{}
	//recipe := objects.Recipe{Id: id_rcp}

	//err := this.db.Model(&recipe).Association("Ingredients").Find(&temp).Error
	//err := this.db.Preload("Ingredients").Preload("Recipes").Find(&temp).Error
	err := this.db.Where(objects.RecipeIngredient{Recipe: idRecipe}).Find(&temp).Error
	return temp, err
}