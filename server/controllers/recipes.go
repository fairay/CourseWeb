package controllers

import (
	"api/recipes/controllers/responses"
	auth "api/recipes/controllers/token"
	"api/recipes/errors"
	"api/recipes/models"
	"api/recipes/objects"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type recipe struct {
	model *models.RecipeM
}

func InitRecipes(r *mux.Router, model *models.RecipeM) {
	ctrl := &recipe{model}
	r.HandleFunc("/accounts/{login}/recipes", ctrl.getRecipesByLogin).Methods("GET")

	r.HandleFunc("/recipes", ctrl.getAllRecipes).Methods("GET")
	r.HandleFunc("/recipes", ctrl.addRecipe).Methods("POST")
	r.HandleFunc("/recipes/{id}", ctrl.deleteRecipe).Methods("DELETE")
	r.HandleFunc("/recipes/{id}/like", ctrl.deleteRecipe).Methods("PUT")
}

// @Tags Recipes
// @Router /recipes [get]
// @Summary Retrieves all recipes
// @Produce json
// @Success 200 {object} []objects.RecipeDTO
func (this *recipe) getAllRecipes(w http.ResponseWriter, r *http.Request) {
	data := this.model.GetAll()
	responses.JsonSuccess(w, objects.Recipe{}.ArrToDTO(data))
}

// @Tags Recipes
// @Router /accounts/{login}/recipes [get]
// @Summary Retrieves recipes of user
// @Param login path string true "Category title"
// @Produce json
// @Success 200 {object} []objects.RecipeDTO
func (this *recipe) getRecipesByLogin(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	login := urlParams["login"]
	data, _ := this.model.FindByLogin(login)
	responses.JsonSuccess(w, objects.Recipe{}.ArrToDTO(data))
}

// @Tags Recipes
// @Router /recipes [post]
// @Param recipe body objects.RecipeDTO true "Recipe data"
// @Summary Creation a new recipe
// @Produce json
// @Success 201 {object} objects.RecipeDTO
func (this *recipe) addRecipe(w http.ResponseWriter, r *http.Request) {
	rcpDTO := new(objects.RecipeDTO)
	err := json.NewDecoder(r.Body).Decode(rcpDTO)
	if err != nil {
		responses.BadRequest(w, "Invalid request")
		return
	}

	login, err := auth.LoginFromCookie(r)
	if err != nil {
		responses.AuthenticationFailed(w)
		return
	} else {
		rcpDTO.Author = login
	}

	err = this.model.AddRecipe(rcpDTO.ToModel())
}

// @Tags Recipes
// @Router /recipes/{id} [delete]
// @Param id path int true "Recipe's id"
// @Summary Deletion the recipe by its id
// @Success 200
func (this *recipe) deleteRecipe(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]

	login, err := auth.LoginFromCookie(r)
	if err != nil {
		responses.AuthenticationFailed(w)
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Invalid request")
	}

	err = this.model.DeleteRecipe(id, login)
	switch err {
	case nil:
		responses.TextSuccess(w, "Deletion was successful")
	case errors.AccessDeleteDenied:
		responses.AccessDenied(w)
	case errors.RecordNotFound:
		responses.RecordNotFound(w, "user")
	case errors.UnknownAccount:
		responses.RecordNotFound(w, "author")
	case errors.UnknownRecipe:
		responses.RecordNotFound(w, "recipe")
	}
}

