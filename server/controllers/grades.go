package controllers

import (
	"api/recipes/controllers/responses"
	auth "api/recipes/controllers/token"
	"api/recipes/errors"
	"api/recipes/models"
	"api/recipes/objects"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type likesCtrl struct {
	recM *models.RecipeM
	accM *models.AccountM
}

func InitLikes(r *mux.Router, recM *models.RecipeM, accM *models.AccountM) {
	ctrl := &likesCtrl{recM, accM}

	r.HandleFunc("/recipes/{id}/like", ctrl.add).Methods("PUT")
	r.HandleFunc("/recipes/{id}/like", ctrl.del).Methods("DELETE")
	r.HandleFunc("/recipes/{id}/like", ctrl.getByRecipe).Methods("GET")
	r.HandleFunc("/accounts/{login}/like", ctrl.getByUser).Methods("GET")
}

// @Tags Likes
// @Router /recipes/{id}/like [put]
// @Summary Adds like to recipe from authorized user
// @Param id path int true "Recipe id"
// @Produce json
// @Success 201 Successful operation
// @Failure 400 Invalid value
// @Failure 401 Authentication failed
func (this *likesCtrl) add(w http.ResponseWriter, r *http.Request) {
	login, err := auth.LoginFromCookie(r)
	if err != nil {
		responses.AuthenticationFailed(w)
		return
	}

	urlParams := mux.Vars(r)
	strId := urlParams["id"]

	id_rcp, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Wrong recipe's id")
		return
	}

	err = this.recM.AddGrade(id_rcp, login)
	switch err {
	case nil:
		responses.TextSuccess(w, "The recipe was successfully liked")
	case errors.UnknownRecipe:
		responses.RecordNotFound(w, "recipe")
	case errors.UnknownAccount:
		responses.RecordNotFound(w, "user")
	default:
		responses.BadRequest(w, "Failed to like this recipe")
	}
}

// TODO:
// @Tags Likes
// @Router /recipes/{id}/like [delete]
// @Summary Deletes like to recipe from authorized user
// @Param id path int true "Recipe id"
// @Produce json
// @Success 200
func (this *likesCtrl) del(w http.ResponseWriter, r *http.Request) {
	fmt.Print("del")
}

// TODO:
// @Tags Likes
// @Router /recipes/{id}/like [get]
// @Summary Retrieves all liked users
// @Param id path int true "Recipe id"
// @Produce json
// @Success 200 {object} []objects.AccountDTO
func (this *likesCtrl) getByRecipe(w http.ResponseWriter, r *http.Request) {
	fmt.Print("getByRecipe")
}

// @Tags Likes
// @Router /accounts/{login}/like [get]
// @Summary Retrieves all user's liked recipes
// @Param login path string true "Requested account"
// @Produce json
// @Success 200 {object} objects.RecipeDTO
// @Failure 400 Invalid value
func (this *likesCtrl) getByUser(w http.ResponseWriter, r *http.Request) {
	var login string
	var err error

	urlParams := mux.Vars(r)
	login = urlParams["login"]

	data, err := this.recM.GetLikedByLogin(login)
	switch err {
	case nil:
		responses.JsonSuccess(w, objects.Recipe{}.ArrToDTO(data))
	case errors.UnknownAccount:
		responses.RecordNotFound(w, "user")
	default:
		responses.BadRequest(w, "Error in getting liked recipes")
	}
}
