package repository

import (
	"api/recipes/objects"

	"github.com/jinzhu/gorm"
)

type RecipesRep interface {
	List() ([]objects.Recipe)
	Create(rcp objects.Recipe) error
}

type PGRecipesRep struct {
	db *gorm.DB
}

func NewRecipesRep (db *gorm.DB) *PGRecipesRep {
	return &PGRecipesRep{db}
}

func (this *PGRecipesRep) List() ([]objects.Recipe) {
	temp := []objects.Recipe{}
	this.db.Find(&temp)
	return temp
}

func (this *PGRecipesRep) Create(rcp objects.Recipe) error {
	return this.db.Create(&rcp).Error
}