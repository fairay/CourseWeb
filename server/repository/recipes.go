package repository

import (
	"api/recipes/objects"

	"github.com/jinzhu/gorm"
)

type RecipesRep interface {
	List() []objects.Recipe
	FindByLogin(login string) []objects.Recipe
	FindById(id int) (objects.Recipe, error)
	Create(rcp *objects.Recipe) error
	Delete(id int) error
}

type PGRecipesRep struct {
	db *gorm.DB
}

func NewRecipesRep(db *gorm.DB) *PGRecipesRep {
	return &PGRecipesRep{db}
}

func (this *PGRecipesRep) List() []objects.Recipe {
	temp := []objects.Recipe{}
	this.db.Find(&temp)
	return temp
}

func (this *PGRecipesRep) FindByLogin(login string) []objects.Recipe {
	temp := []objects.Recipe{}
	this.db.Where("author = ?", login).Find(&temp)
	return temp
}

func (this *PGRecipesRep) FindById(id int) (objects.Recipe, error) {
	temp := objects.Recipe{}
	err := this.db.Where("id = ?", id).Find(&temp).Error
	return temp, err
}

func (this *PGRecipesRep) Create(obj *objects.Recipe) error {
	obj.Id = 0
	return this.db.Create(obj).Error
}

func (this *PGRecipesRep) Delete(id int) error {
	recipe := objects.Recipe{Id: id}

	err := this.db.Model(&recipe).Association("Categories").Clear().Error
	if err != nil { return err }
	
	return this.db.Where("id = ?", id).Delete(&objects.Recipe{}).Error
}
