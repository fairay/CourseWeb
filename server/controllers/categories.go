package controllers

import (
	"api/recipes/repository"
	"api/recipes/repository/objects"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type category struct {
	db *gorm.DB
}

func InitCategories(r *mux.Router, db *gorm.DB) {
	ctrl := &category{db}
	r.HandleFunc("/categories", ctrl.getAllCategories).Methods("GET")
}

// @Tags Categories
// @Router /categories [get]
// @Summary Retrieves all categories
// @Produce json
// @Success 200 {object} []objects.Categories
func (this *category) getAllCategories(w http.ResponseWriter, r *http.Request) {
	rep := &repository.PGCategoriesRep{this.db}
	temp := rep.List()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(temp)
}
