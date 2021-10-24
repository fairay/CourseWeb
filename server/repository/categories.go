package repository

import (
	"api/recipes/objects"
	"strings"

	"github.com/jinzhu/gorm"
)

type CategoriesRep interface {
	List() ([]objects.Categories)
	Find(ctg string) ([]objects.Categories)
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

func (this *PGCategoriesRep) Find(ctg string) ([]objects.Categories) { 
	temp := []objects.Categories{}
	this.db.Where("LOWER(title) LIKE ?", "%" + strings.ToLower(ctg) + "%").Find(&temp)
	return temp
}