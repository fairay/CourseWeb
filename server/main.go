package main

import (
	"app/server/repository/objects"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

/*var e error*/

func main() {
	dsn := "host=localhost user=admin password=admin dbname=recipes port=5432 sslmode=disable"
	db, e := gorm.Open("postgres", dsn)

	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Connection Established")
	}

	defer db.Close()
	db.SingularTable(true)
	db.AutoMigrate(&objects.Categories{})

	fmt.Println("Hello test")
	GetAllCategories()

	router := mux.NewRouter()
	router.HandleFunc("/test", getTest).Methods("GET")
	//router.HandleFunc("/categories", GetAllCategories).Methods("GET")
	http.ListenAndServe(":1991", router)
}

func getTest(w http.ResponseWriter, r *http.Request) {
	var data = [...]int{155555, 2, 3, 4}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(data)
}

func GetAllCategories() {
	temp := []objects.Categories{}
	db.Find(&temp)
	fmt.Println(temp)

	/*w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(temp)*/
}
