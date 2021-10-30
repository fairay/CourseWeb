package models

import (
	//"api/recipes/objects"
	"api/recipes/objects"
	"api/recipes/repository"
	"api/recipes/errors"
)

type AccountM struct {
	rep repository.AccountsRep
	models *Models
}

func NewAccount(rep repository.AccountsRep, models *Models) *AccountM {
	return &AccountM{rep, models}
}

/*func (this *AccountM) GetAll() ([]objects.Account) {
	temp := this.rep.List()
	return temp
}

func (this *AccountM) Get(ctg string) (objects.Account) {
	temp := this.rep.Get(ctg)
	return temp
}*/

func (this *AccountM) find(login string) (*objects.Account, error) {
	return this.rep.Find(login)
}

func (this *AccountM) IsExists(login string) bool {
	_, err := this.find(login)

	if err != nil {
		return false
	}

	return true
}

func (this *AccountM) LogIn(login string, password string) (*objects.Account, error){
	acc, err := this.find(login)

	if err != nil {
		return nil, err
	}

	// TODO: change the way of compare 
	if password != acc.HashedPassword {
		return nil, errors.WrongPassword
	}

	return acc, err
}
