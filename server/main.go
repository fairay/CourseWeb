package main

import (
	"api/recipes/controllers"
	auth "api/recipes/controllers/token"
	_ "api/recipes/docs"
	"api/recipes/models"
	_ "api/recipes/objects"
	"api/recipes/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initDBConnection(cnf utils.DBConfiguration) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cnf.Host, cnf.User, cnf.Password, cnf.Name, cnf.Port)
	db, e := gorm.Open(cnf.Type, dsn)

	if e != nil {
		utils.Logger.Print("DB Connection failed")
		utils.Logger.Fatal(e)
	} else {
		utils.Logger.Print("DB Connection Established")
	}

	// TODO: learn db connection setup actions
	db.SingularTable(true)
	return db
}

func initControllers(r *mux.Router, m *models.Models) {
	r.Use(utils.LogHandler)
	r.Use(auth.JwtAuthentication)

	controllers.InitCategories(r, m.Category)
	controllers.InitRecipes(r, m.Recipes)
	controllers.InitAccount(r, m.Accounts)
	controllers.InitLikes(r, m.Recipes, m.Accounts)
}

// @title Recipes API
// @version 0.1
// @description API for culinary recipes (BMSTU Web development cource project)
func main() {
	utils.InitConfig()
	utils.InitLogger()
	defer utils.CloseLogger()

	db := initDBConnection(utils.Config.DB)
	defer db.Close()

	router := mux.NewRouter()
	models := models.InitModels(db)
	initControllers(router, models)

	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	utils.Logger.Print("Server started")
	fmt.Printf("Server is running on http://localhost:%d\n", utils.Config.Port)
	code := http.ListenAndServe(fmt.Sprintf(":%d", utils.Config.Port), router)

	utils.Logger.Printf("Server ended with code %s", code)
}
