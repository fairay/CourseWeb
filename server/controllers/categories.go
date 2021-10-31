package controllers

import (
	"api/recipes/controllers/responses"
	"api/recipes/models"
	"api/recipes/objects"
	"net/http"

	"github.com/gorilla/mux"
)

type category struct {
	model *models.CategoryM
}

func InitCategories(r *mux.Router, model *models.CategoryM) {
	ctrl := &category{model}
	r.HandleFunc("/categories", ctrl.getAllCategories).Methods("GET")
	r.HandleFunc("/categories/{title}", ctrl.getCategory).Methods("GET")
}

// @Tags Categories
// @Router /categories [get]
// @Summary Retrieves all categories
// @Param search query string false "Search query"
// @Produce json
// @Success 200 {object} []objects.CategoryDTO
func (this *category) getAllCategories(w http.ResponseWriter, r *http.Request) {
	urlParams := r.URL.Query()
	search := urlParams.Get("search")
	data := this.model.Find(search)
	responses.JsonSuccess(w, objects.Category{}.ArrToDTO(data))
}

// @Tags Categories
// @Router /categories/{title} [get]
// @Summary Retrieves categoriey
// @Param title path string true "Category title"
// @Produce json
// @Success 200 {object} objects.CategoryDTO
func (this *category) getCategory(w http.ResponseWriter, r *http.Request) {
	temp := mux.Vars(r)
	data := this.model.Get(temp["title"])
	responses.JsonSuccess(w, data.ToDTO())
}
