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

func (this *CategoryM) GetAll() ([]objects.Categories) {
	temp := this.rep.List()
	return temp
}
