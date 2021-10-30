package models

import (
	"api/recipes/objects"
	"api/recipes/repository"
	"api/recipes/errors"
)

type RecipeM struct {
	rep    repository.RecipesRep
	models *Models
}

func NewRecipe(rep repository.RecipesRep, models *Models) *RecipeM {
	return &RecipeM{rep, models}
}

func (this *RecipeM) GetAll() []objects.Recipe {
	temp := this.rep.List()
	return temp
}

func (this *RecipeM) GetAuthor(id int) (*objects.Account, error) {
	rcp, err := this.FindById(id)
	if err != nil {
		return nil, err
	}

	login := rcp.Author

	return this.models.Accounts.find(login)
}

func (this *RecipeM) FindByLogin(login string) ([]objects.Recipe, error) {
	isExist := this.models.Accounts.IsExists(login)
	if isExist == false {
		return nil, nil
	}

	temp := this.rep.FindByLogin(login)
	return temp, nil
}

func (this *RecipeM) FindById(id int) (*objects.Recipe, error) {
	rcp, err := this.rep.FindById(id)
	if err != nil {
		return nil, err
	}

	return &rcp, err
}

func (this *RecipeM) AddRecipe(obj *objects.Recipe) (err error) {
	_, err = this.models.Accounts.find(obj.Author)
	if err != nil {
		return err
	}

	// TODO: validate other data
	err = this.rep.Create(obj)
	return err
}

func (this *RecipeM) DeleteRecipe(id int, login string) (err error) {
	userRole, err := this.models.Accounts.GetRole(login)
	if err != nil {
		return err
	}

	author, err := this.GetAuthor(id)
	if err != nil {
		return err
	}

	authorRole := author.Role

	if userRole == AdminRole || userRole == authorRole {
		this.rep.Delete(id)
		err = nil
	}else {
		err = errors.AccessDeleteDenied
	}

	return err
}
