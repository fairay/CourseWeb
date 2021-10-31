package repository

import (
	"api/recipes/objects"
	"strings"

	"github.com/jinzhu/gorm"
)

type CategoriesRep interface {
	List() ([]objects.Category)
	Find(ctg string) ([]objects.Category, error)
	FindRecipes(ctg string) ([]objects.Recipe, error)

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

func (this *PGCategoriesRep) Find(ctg string) ([]objects.Category, error) { 
	temp := []objects.Category{}
	err := this.db.Where("LOWER(title) LIKE ?", "%" + strings.ToLower(ctg) + "%").Find(&temp).Error
	return temp, err
}

func (this *PGCategoriesRep) FindRecipes(ctg string) ([]objects.Recipe, error) { 
	temp := []objects.Recipe{}
	category := objects.Category{Title: ctg}

	err := this.db.Model(&category).Association("Recipes").Find(&temp).Error

	return temp, err
}

func (this *PGCategoriesRep) Get(ctg string) (objects.Category) { 
	temp := objects.Category{}
	this.db.Where("LOWER(title) = ?", strings.ToLower(ctg)).Find(&temp)
	return temp
}