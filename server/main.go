package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main()  {
	fmt.Println("Hello test")

	router := mux.NewRouter()
	router.HandleFunc("/test", getTest).Methods("GET")
	http.ListenAndServe(":1991", router)
}

func getTest(w http.ResponseWriter, r *http.Request) {
	var data = [...]int{1, 2, 3, 4}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(data)
}
