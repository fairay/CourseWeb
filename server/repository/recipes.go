package repository

import (
	"api/recipes/objects"

	"github.com/jinzhu/gorm"
)

type RecipesRep interface {
	List() ([]objects.Recipes)
}

type PGRecipesRep struct {
	db *gorm.DB
}

func NewRecipesRep (db *gorm.DB) *PGRecipesRep {
	return &PGRecipesRep{db}
}

func (this *PGRecipesRep) List() ([]objects.Recipes) {
	temp := []objects.Recipes{}
	this.db.Find(&temp)
	return temp
}