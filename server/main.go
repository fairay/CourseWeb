package main

import (
	"api/recipes/controllers"
	_ "api/recipes/docs"
	_ "api/recipes/objects"
	"api/recipes/utils"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	db.SingularTable(true)
	return db
}
/*
func initControllers(r *mux.Router, m *models.Models) {
	r.Use(utils.LogHandler)
	r.Use(auth.JwtAuthentication)

	controllers.InitCategories(r, m.Category)
	controllers.InitRecipes(r, m.Recipes)
	controllers.InitAccount(r, m.Accounts)
	controllers.InitIngredients(r, m.Ingredients)
	controllers.InitLikes(r, m.Recipes, m.Accounts)
	controllers.InitSteps(r, m.Steps)
}

func InitRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	models := models.InitModels(db)
	initControllers(router, models)
	return router
}

func runSwagger(r *mux.Router) {
	r.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
}

func RunRouter(r *mux.Router, port uint16) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
*/

// @Title Recipes API
// @Version 0.1
// @Description API for culinary recipes (BMSTU Web development cource project)
// @securityDefinitions.basic BasicAuth
func main() {
	rand.Seed(time.Now().UnixNano())

	utils.InitConfig(os.Args)
	utils.InitLogger()
	defer utils.CloseLogger()

	db := initDBConnection(utils.Config.DB)
	defer db.Close()

	r := controllers.InitRouter(db)
	controllers.RunSwagger(r);
	utils.Logger.Print("Server started")
	fmt.Printf("Server is running on http://localhost:%d\n", utils.Config.Port)
	code := controllers.RunRouter(r, utils.Config.Port);

	utils.Logger.Printf("Server ended with code %s", code)
}
