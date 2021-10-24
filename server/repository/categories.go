package repository

import (
	"api/recipes/repository/objects"

	"github.com/jinzhu/gorm"
)

type CategoriesRep interface {
	List() ([]objects.Categories)
}

type PGCategoriesRep struct {
	db *gorm.DB
}

func (this *PGCategoriesRep) List() ([]objects.Categories) {
	temp := []objects.Categories{}
	this.db.Find(&temp)
	return temp
}
