package tests

import (
	"api/recipes/controllers"
	"api/recipes/objects"
	"api/recipes/utils"

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
	port = 9000
	path := "9000.txt"

	//utils.InitConfig(strconv.Itoa(int(port)))
	utils.InitLogger(path)
	
	db, err := StubConnecton()
	if err != nil {	panic(err) }

	r := controllers.InitRouter(db)

	go controllers.RunRouter(r, port);
	return port
}
