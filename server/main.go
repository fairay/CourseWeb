package main

import (
	"api/recipes/controllers"
	_ "api/recipes/docs"
	"api/recipes/models"
	_ "api/recipes/objects"
	"api/recipes/repository"
	"api/recipes/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initDBConnection(cnf utils.DBConfiguration) (*gorm.DB) {
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
	// TODO: need db.AutoMigrate(&objects.Categories{}) ?
	db.SingularTable(true)
	return db
}

func initControllers(r *mux.Router, db *gorm.DB) {
	controllers.InitCategories(r, models.NewCategory(repository.NewCategotiesRep(db)))
	controllers.InitRecipes(r, models.NewRecipe(repository.NewRecipesRep(db)))
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
	initControllers(router, db)

	router.HandleFunc("/test", getTest).Methods("GET")
	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	utils.Logger.Print("Server started")
	fmt.Printf("Server is running on http://localhost%s\n", utils.Config.Port)
	http.ListenAndServe(utils.Config.Port, router)
}

func getTest(w http.ResponseWriter, r *http.Request) {
	var data = [...]int{155555, 2, 3, 4}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(data)
}
