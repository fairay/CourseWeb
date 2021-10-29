package repository

import (
	"api/recipes/objects"
	"api/recipes/errors"

	"github.com/jinzhu/gorm"
)

type AccountsRep interface {
	Find(login string) (*objects.Account, error)
}

type PGAccountsRep struct {
	db *gorm.DB
}

func NewAccountsRep (db *gorm.DB) *PGAccountsRep {
	return &PGAccountsRep{db}
}

func (this *PGAccountsRep) Find(login string) (*objects.Account, error) { 
	temp := new (objects.Account)
	err := this.db.Where("login = ?", login).First(temp).Error
	switch err {
	case nil:
		break
	case gorm.ErrRecordNotFound:
		err = errors.RecordNotFound
		break
	default:
		err = errors.UnknownError
	}

	return temp, err
}