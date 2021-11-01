package controllers

import (
	"api/recipes/controllers/responses"
	"api/recipes/errors"
	"api/recipes/models"
	"api/recipes/objects"
	"strconv"

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
	data, _ := this.model.Find(search)
	responses.JsonSuccess(w, objects.Category{}.ArrToDTO(data))
}

// @Tags Categories
// @Router /categories/{title} [get]
// @Summary Retrieves category by title
// @Param title path string true "Category title"
// @Produce json
// @Success 200 {object} objects.CategoryDTO
func (this *category) get(w http.ResponseWriter, r *http.Request) {
	temp := mux.Vars(r)
	data := this.model.Get(temp["title"])
	responses.JsonSuccess(w, data.ToDTO())
}

// @Tags Categories
// @Router /categories/{title}/recipes [get]
// @Summary Retrieves all recipes at this category
// @Param title path string true "Searched category"
// @Produce json
// @Success 200 {object} []objects.RecipeDTO
func (this *category) getRecipes(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	ctg := urlParams["title"]

	data, err := this.model.GetRecipes(ctg)
	switch err {
	case nil:
		responses.JsonSuccess(w, objects.Recipe{}.ArrToDTO(data))
	case errors.UnknownRecipe:
		responses.RecordNotFound(w, "category")
	default:
		responses.BadRequest(w, "Error in getting recipes")
	}
}

// @Tags Categories
// @Router /recipes/{id}/categories [get]
// @Summary Retrieves all recipe's categories
// @Param id path int true "Recipe's id"
// @Produce json
// @Success 200 {object} []objects.CategoryDTO
func (this *category) getByRecipe(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]

	id_rcp, err := strconv.Atoi(strId)

	data, err := this.model.GetByRecipe(id_rcp)

	switch err {
	case nil:
		responses.JsonSuccess(w, objects.Category{}.ArrToDTO(data))
	case errors.UnknownRecipe:
		responses.RecordNotFound(w, "recipe")
	default:
		responses.BadRequest(w, "Error in getting categories")
	}
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
// @Summary Adds category
// @Param id path int true "Recipe's id"
// @Param recipe body objects.CategoryDTO true "Category"
// @Produce json
// @Success 200
func (this *category) putToRecipe(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Categories
// @Router /recipes/{id}/categories [delete]
// @Summary Removes category
// @Param id path int true "Recipe's id"
// @Param recipe body objects.CategoryDTO true "Category"
// @Produce json
// @Success 200
func (this *category) delFromRecipe(w http.ResponseWriter, r *http.Request) {
}
