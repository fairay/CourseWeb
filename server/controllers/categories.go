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
	r.HandleFunc("/categories", ctrl.getAll).Methods("GET")
	r.HandleFunc("/categories/{title}", ctrl.get).Methods("GET")
	r.HandleFunc("/categories/{title}/recipes", ctrl.getRecipes).Methods("GET")

	r.HandleFunc("/recipes/{id}/categories", ctrl.getByRecipe).Methods("GET")
	r.HandleFunc("/recipes/{id}/categories", ctrl.putToRecipe).Methods("PUT")
	r.HandleFunc("/recipes/{id}/categories", ctrl.postToRecipe).Methods("POST")
	r.HandleFunc("/recipes/{id}/categories", ctrl.delFromRecipe).Methods("DELETE")
}

// @Tags Categories
// @Router /categories [get]
// @Summary Retrieves all categories
// @Param search query string false "Search query"
// @Produce json
// @Success 200 {object} []objects.CategoryDTO
func (this *category) getAll(w http.ResponseWriter, r *http.Request) {
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
func (this *category) get(w http.ResponseWriter, r *http.Request) {
	temp := mux.Vars(r)
	data := this.model.Get(temp["title"])
	responses.JsonSuccess(w, data.ToDTO())
}

// TODO:
// @Tags Categories
// @Router /categories/{title}/recipes [get]
// @Summary Retrieves all recipes with such category
// @Param title path string true "Searched category"
// @Produce json
// @Success 200 {object} []objects.RecipeDTO
func (this *category) getRecipes(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Categories
// @Router /recipes/{id}/categories [get]
// @Summary Retrieves all categories
// @Param id path int true "Recipe's id"
// @Produce json
// @Success 200 {object} []objects.CategoryDTO
func (this *category) getByRecipe(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Categories
// @Router /recipes/{id}/categories [post]
// @Summary Posts all categories
// @Param id path int true "Recipe's id"
// @Param categories body []objects.CategoryDTO true "Categories"
// @Produce json
// @Success 201 {object} []objects.CategoryDTO
func (this *category) postToRecipe(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Categories
// @Router /recipes/{id}/categories [put]
// @Summary Adding category
// @Param id path int true "Recipe's id"
// @Param recipe body objects.CategoryDTO true "Category"
// @Produce json
// @Success 200
func (this *category) putToRecipe(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Categories
// @Router /recipes/{id}/categories [delete]
// @Summary Deleting category
// @Param id path int true "Recipe's id"
// @Param recipe body objects.CategoryDTO true "Category"
// @Produce json
// @Success 200
func (this *category) delFromRecipe(w http.ResponseWriter, r *http.Request) {
}
