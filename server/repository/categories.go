package repository

import (
	"api/recipes/objects"

	"github.com/jinzhu/gorm"
)

type CategoriesRep interface {
	List() ([]objects.Categories)
}

type PGCategoriesRep struct {
	db *gorm.DB
}

func NewCategotiesRep (db *gorm.DB) *PGCategoriesRep {
	return &PGCategoriesRep{db}
}

func (this *PGCategoriesRep) List() ([]objects.Categories) {
	temp := []objects.Categories{}
	this.db.Find(&temp)
	return temp
}
