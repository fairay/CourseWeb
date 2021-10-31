package repository

import (
	"api/recipes/objects"

	"github.com/jinzhu/gorm"
)

type StepsRep interface {
	List(recipeId int) ([]objects.Step, error)
}

type PGStepsRep struct {
	db *gorm.DB
}

func NewStepsRep(db *gorm.DB) *PGStepsRep {
	return &PGStepsRep{db}
}

func (this *PGStepsRep) List(recipeId int) (arr []objects.Step, err error) {
	err = this.db.Where("recipe = ?", recipeId).Find(&arr).Error
	return
}
