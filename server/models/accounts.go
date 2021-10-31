package models

import (
	"api/recipes/objects"
	"api/recipes/repository"
	"api/recipes/errors"
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

func (this *AccountM) UpdateRole(login, role string) error {
	if this.IsExists(login) == false {
		return errors.UnknownAccount
	}

	if role != AdminRole && role != UserRole {
		return errors.UnknownRole
	}

	// FIXME: current role == admin

	return this.rep.UpdateRole(login, role)
}

func (this *AccountM) Find(login string) (*objects.Account, error) {
	return this.rep.Find(login)
}

func (this *AccountM) IsExists(login string) bool {
	data, _ := this.Find(login)

	if data == nil { return false }

	return true
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
