package repository

import (
	"api/recipes/objects"

	"github.com/jinzhu/gorm"
)

type StepsRep interface {
	List(recipeId int) ([]objects.Step, error)
	FindSteps(id_rcp int) ([]objects.Step, error)
	FindStepByNum(id_rcp, step int) (objects.Step, error)
	Create(obj *objects.Step) error
	Delete(id_rcp, step int) error
	UpdateStep(id_rcp int, step int, obj *objects.Step) error
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

func (this *PGStepsRep) FindSteps(id_rcp int) ([]objects.Step, error) {
	temp := []objects.Step{}
	err := this.db.Where("recipe = ?", id_rcp).Order("num asc").Find(&temp).Error

	return temp, err
}

func (this *PGStepsRep) FindStepByNum(id_rcp, step int) (objects.Step, error) {
	temp := objects.Step{}
	err := this.db.Where("recipe = ? AND num = ?", id_rcp, step).Find(&temp).Error

	return temp, err
}

func (this *PGStepsRep) Create(obj *objects.Step) error {
	return this.db.Create(obj).Error
}

func (this *PGStepsRep) Delete(id_rcp, step int) error {
	temp := objects.Step{}
	return this.db.Where("recipe = ? AND num = ?", id_rcp, step).Delete(&temp).Error
}

func (this *PGStepsRep) UpdateStep(id_rcp int, step int, obj *objects.Step) error {
	return this.db.Model(&objects.Step{}).Where("recipe = ? AND num = ?", id_rcp, step).Update(objects.Step{Description: obj.Description, Title: obj.Title}).Error
}