package main

import (
	_ "api/recipes/docs"
	"api/recipes/repository/objects"
	"api/recipes/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Handler struct {
	db *gorm.DB
}

// @title Recipes API
// @version 0.1
// @description API for culinary recipes (BMSTU Web development cource project)

func main() {
	utils.InitConfig()
	utils.InitLogger()
	defer utils.CloseLogger()

	cnf := utils.Config.DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cnf.Host, cnf.User, cnf.Password, cnf.Name, cnf.Port)
	db, e := gorm.Open(cnf.Type, dsn)

	if e != nil {
		fmt.Println(e)
	} else {
		utils.Logger.Print("Connection Established")
	}

	defer db.Close()
	db.SingularTable(true)
	db.AutoMigrate(&objects.Categories{})

	handler := new(Handler)
	handler.db = db

	router := mux.NewRouter()
	router.HandleFunc("/test", getTest).Methods("GET")
	router.HandleFunc("/categories", handler.getAllCategories).Methods("GET")
	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	utils.Logger.Print("Server started")
	http.ListenAndServe(utils.Config.Port, router)
}

func getTest(w http.ResponseWriter, r *http.Request) {
	var data = [...]int{155555, 2, 3, 4}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(data)
}

// @Tags Categories
// @Router /categories [get]
// @Summary Retrieves all categories
// @Produce json
// @Success 200 {object} []objects.Categories
func (handler *Handler) getAllCategories(w http.ResponseWriter, r *http.Request) {
	temp := []objects.Categories{}
	handler.db.Find(&temp)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(temp)
}
