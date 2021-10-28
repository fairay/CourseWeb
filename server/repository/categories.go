package repository

import (
	"api/recipes/objects"
	"strings"

	"github.com/jinzhu/gorm"
)

type CategoriesRep interface {
	List() ([]objects.Category)
	Find(ctg string) ([]objects.Category)

	Get(ctg string) (objects.Category)
}

type PGCategoriesRep struct {
	db *gorm.DB
}

func NewCategotiesRep (db *gorm.DB) *PGCategoriesRep {
	return &PGCategoriesRep{db}
}

func (this *PGCategoriesRep) List() ([]objects.Category) {
	temp := []objects.Category{}
	this.db.Find(&temp)
	return temp
}

func (this *PGCategoriesRep) Find(ctg string) ([]objects.Category) { 
	temp := []objects.Category{}
	this.db.Where("LOWER(title) LIKE ?", "%" + strings.ToLower(ctg) + "%").Find(&temp)
	return temp
}

func (this *PGCategoriesRep) Get(ctg string) (objects.Category) { 
	temp := objects.Category{}
	this.db.Where("LOWER(title) = ?", strings.ToLower(ctg)).Find(&temp)
	return temp
}