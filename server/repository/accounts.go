package repository

import (
	"api/recipes/errors"
	"api/recipes/objects"

	"github.com/jinzhu/gorm"
)

type AccountsRep interface {
	Create(obj *objects.Account) error
	UpdateRole(login, role string) error
	Find(login string) (*objects.Account, error)
	FindLikedRecipe(id_rcp int) ([]objects.Account, error)
}

type PGAccountsRep struct {
	db *gorm.DB
}

func NewAccountsRep(db *gorm.DB) *PGAccountsRep {
	return &PGAccountsRep{db}
}

func (this *PGAccountsRep) Create(obj *objects.Account) error {
	return this.db.Create(obj).Error
}

func (this *PGAccountsRep) UpdateRole(login, role string) error {
	return this.db.Model(&objects.Account{}).Where("login = ?", login).Update("role", role).Error
}

func (this *PGAccountsRep) Find(login string) (*objects.Account, error) {
	temp := new(objects.Account)
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

func (this *PGAccountsRep) FindLikedRecipe(id_rcp int) ([]objects.Account, error) {
	temp := []objects.Account{}
	recipe := objects.Recipe{Id: id_rcp}

	err := this.db.Model(&recipe).Association("Accounts").Find(&temp).Error

	return temp, err
}