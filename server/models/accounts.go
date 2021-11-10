package models

import (
	"api/recipes/errors"
	"api/recipes/objects"
	"api/recipes/repository"
)

const (
	AdminRole string = "admin"
	UserRole string = "user"
	UnauthRole string = "unauthorized"
)

type AccountM struct {
	rep repository.AccountsRep
	models *Models
}

func NewAccount(rep repository.AccountsRep, models *Models) *AccountM {
	return &AccountM{rep, models}
}

func (this *AccountM) Create(obj *objects.Account) error {
	if this.IsExists(obj.Login) {
		return errors.AccountExists
	}

	obj.Role = UserRole
	obj.Salt = ""

	err := this.rep.Create(obj)
	if err != nil {
		return errors.DBAdditionError
	}

	return err
}

func (this *AccountM) UpdateRole(cur_login, login, role string) error {
	if this.models.Accounts.IsExists(cur_login) == false {
		return errors.UnknownAccount
	}

	cur_role, err := this.GetRole(cur_login)
	if err != nil {
		return errors.UnknownAccount
	}

	if cur_role != AdminRole {
		return errors.AccessDenied
	}

	if this.IsExists(login) == false {
		return errors.UnknownAccount
	}

	if role != AdminRole && role != UserRole {
		return errors.UnknownRole
	}

	return this.rep.UpdateRole(login, role)
}

func (this *AccountM) Find(login string) (*objects.Account, error) {
	return this.rep.Find(login)
}

func (this *AccountM) FindLikedRecipe(id_rcp int) ([]objects.Account, error) {
	_, err := this.models.Recipes.FindById(id_rcp)
	if err != nil {
		return nil, errors.UnknownRecipe
	}

	return this.rep.FindLikedRecipe(id_rcp)
}

func (this *AccountM) IsExists(login string) bool {
	_, err := this.Find(login)

	return err == nil
}

func (this *AccountM) GetRole(login string) (role string, err error) {
	acc, err := this.Find(login)

	if err != nil {
		return "", err
	}

	return acc.Role, err
}

func (this *AccountM) LogIn(login string, password string) (*objects.Account, error){
	acc, err := this.Find(login)

	if err != nil {
		return nil, err
	}

	// TODO: change the way of compare 
	if password != acc.HashedPassword {
		return nil, errors.WrongPassword
	}

	return acc, err
}