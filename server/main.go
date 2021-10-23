package main

import (
	"api/recipes/repository/objects"
	"api/recipes/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Handler struct {
	db *gorm.DB
}

func main() {
	utils.InitConfig()
	cnf := utils.Config.DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		cnf.Host, cnf.User, cnf.Password, cnf.Name, cnf.Port)
	db, e := gorm.Open(cnf.Type, dsn)

	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Connection Established")
		log.Fatal()
	}

	defer db.Close()
	db.SingularTable(true)
	db.AutoMigrate(&objects.Categories{})

	handler := new(Handler)
	handler.db = db

	router := mux.NewRouter()
	router.HandleFunc("/test", getTest).Methods("GET")
	router.HandleFunc("/categories", handler.getAllCategories).Methods("GET")
	http.ListenAndServe(utils.Config.Port, router)
}

func getTest(w http.ResponseWriter, r *http.Request) {
	var data = [...]int{155555, 2, 3, 4}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(data)
}

func (handler *Handler) getAllCategories(w http.ResponseWriter, r *http.Request) {
	temp := []objects.Categories{}
	handler.db.Find(&temp)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(temp)
}
