package controllers

import (
	"api/recipes/models"
	"net/http"

	"github.com/gorilla/mux"
)

type category struct {
	model *models.CategoryM
}

func InitCategories(r *mux.Router, model *models.CategoryM) {
	ctrl := &category{model}
	r.HandleFunc("/categories", ctrl.getAllCategories).Methods("GET")
	r.HandleFunc("/categories/find/{req}", ctrl.getCertainCategories).Methods("GET")
}

// @Tags Categories
// @Router /categories [get]
// @Summary Retrieves all categories
// @Produce json
// @Success 200 {object} []objects.Categories
func (this *category) getAllCategories(w http.ResponseWriter, r *http.Request) {
	data := this.model.GetAll()
	jsonSuccess(w, data)
}

// @Tags Categories
// @Router /categories/find/{req} [get]
// @Summary Retrieves categories with certain item
// @Param req path string false "Search query"
// @Produce json
// @Success 200 {object} []objects.Categories
func (this *category) getCertainCategories(w http.ResponseWriter, r *http.Request) {
	temp := mux.Vars(r)
	data := this.model.Find(temp["req"])
	jsonSuccess(w, data)
}
