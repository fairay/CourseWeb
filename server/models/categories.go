package models

import (
	"api/recipes/objects"
	"api/recipes/repository"
)

type CategoryM struct {
	rep repository.CategoriesRep
	models *Models
}

func NewCategory(rep repository.CategoriesRep, models *Models) *CategoryM {
	return &CategoryM{rep, models}
}

func (this *CategoryM) GetAll() ([]objects.Category) {
	temp := this.rep.List()
	return temp
}

func (this *CategoryM) Find(ctg string) ([]objects.Category) {
	temp := this.rep.Find(ctg)
	return temp
}

func (this *CategoryM) Get(ctg string) (objects.Category) {
	temp := this.rep.Get(ctg)
	return temp
}
