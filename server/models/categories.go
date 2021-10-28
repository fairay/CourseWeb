package models

import (
	"api/recipes/objects"
	"api/recipes/repository"
)

type CategoryM struct {
	rep repository.CategoriesRep
}

func NewCategory(rep repository.CategoriesRep) *CategoryM {
	return &CategoryM{rep}
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
