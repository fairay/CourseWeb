package models_test

import (
	"api/recipes/objects"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

/* Used for non existanse check.
This string shouldn't be used in any test objects.
*/
const nWord = "sheeesh"

func stubConnecton() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", ":memory:?cache=shared")
	db.AutoMigrate(
		objects.Recipe{},
		objects.Account{},
		objects.Step{},
		objects.RecipeCategory{}, objects.Category{},
		objects.Ingredient{}, objects.RecipeIngredient{},
	)
	return db, err
}