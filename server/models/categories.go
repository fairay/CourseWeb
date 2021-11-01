package models

import (
	"api/recipes/errors"
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

func (this *CategoryM) Create(obj *objects.Category) error {
	if this.IsExists(obj.Title) {
		return errors.CategoryExists
	}

	err := this.rep.Create(obj)
	if err != nil {
		return errors.DBAdditionError
	}

	return err
}

func (this *CategoryM) GetAll() ([]objects.Category) {
	temp := this.rep.List()
	return temp
}

func (this *CategoryM) Find(ctg string) ([]objects.Category, error) {
	return this.rep.Find(ctg)
}

func (this *CategoryM) Get(ctg string) (*objects.Category, error) {
	data, err := this.rep.Get(ctg)
	if err != nil {
		return nil, errors.RecordNotFound
	}
	return &data, err
}

func (this *CategoryM) GetRecipes(ctg string) ([]objects.Recipe, error) {
	if this.IsExists(ctg) == false {
		return nil, errors.UnknownCategory
	}

	return this.rep.FindRecipes(ctg)
}

func (this *CategoryM) GetByRecipe(id_rcp int) ([]objects.Category, error) {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	return this.rep.FindByRecipe(id_rcp)
}

func (this *CategoryM) PostToRecipe(id_rcp int, ctg_arr *[]objects.Category) error {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}

	for _, ctg := range *ctg_arr {
		if this.IsExists(ctg.Title) == false {
			err = this.Create(&ctg)
			if err != nil { return err }
		}

		err := this.rep.ReplaceInRecipe(id_rcp, ctg.Title)
		if err != nil {
			return err
		}
	}

	return nil
}

func (this *CategoryM) AddToRecipe(id_rcp int, ctg *objects.Category) error {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}
	
	if this.IsExists(ctg.Title) == false {
		err = this.Create(ctg)
		if err != nil { return err }
	}

	err = this.rep.AddToRecipe(id_rcp, ctg.Title)
	if err != nil {
		return err
	}

	return nil
}

func (this *CategoryM) DelFromRecipe(id_rcp int, ctg *objects.Category) error {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return errors.UnknownRecipe
	}
	
	if this.IsExists(ctg.Title) == false {
		return errors.UnknownCategory
	}

	err = this.rep.DelFromRecipe(id_rcp, ctg.Title)
	if err != nil {
		return err
	}

	return nil
}

func (this *CategoryM) IsExistsLike(ctg string) bool {
	data, _ := this.Find(ctg)

	if len(data) == 0 { return false }

	return true
}

func (this *CategoryM) IsExists(ctg string) bool {
	_, err := this.Get(ctg)

	if err != nil { return false }

	return true
}