package tests

import (
	"api/recipes/objects"
	"api/recipes/controllers"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func StubConnecton() (*gorm.DB, error) {
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

func StubServer() (port uint16) {
	db, err := StubConnecton()
	if err != nil {	panic(err) }

	r := controllers.InitRouter(db)
	port = 9000

	go controllers.RunRouter(r, port);
	return port
}
