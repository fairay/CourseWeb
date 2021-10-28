package models

import (
	//"api/recipes/objects"
	"api/recipes/objects"
	"api/recipes/repository"
	"api/recipes/errors"
)

type AccountM struct {
	rep repository.AccountsRep
}

func NewAccount(rep repository.AccountsRep) *AccountM {
	return &AccountM{rep}
}

/*func (this *AccountM) GetAll() ([]objects.Account) {
	temp := this.rep.List()
	return temp
}

func (this *AccountM) Get(ctg string) (objects.Account) {
	temp := this.rep.Get(ctg)
	return temp
}*/

func (this *AccountM) LogIn(login string, password string) (*objects.Account, error){
	acc, err := this.rep.Find(login)

	if err != nil {
		return nil, err
	}

	// TODO: change the way of compare 
	if password != acc.HashedPassword {
		return nil, errors.WrongPassword
	}

	return &acc, err
}