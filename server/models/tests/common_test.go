package models_test

import (
	"api/recipes/objects"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func stubConnecton() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", ":memory:?cache=shared")
	db.AutoMigrate(objects.Category{}, objects.Recipe{}, objects.Account{})
	return db, err
}