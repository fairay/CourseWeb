package controllers

import (
	"api/recipes/models"
	"net/http"

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
	r.HandleFunc("/accounts/like", ctrl.getByUser).Methods("GET")
}

// TODO:
// @Tags Likes
// @Router /recipes/{id}/like [put]
// @Summary Adds like to recipe from authorized user
// @Param id path int true "Recipe id"
// @Produce json
// @Success 201
func (this *likesCtrl) add(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Likes
// @Router /recipes/{id}/like [delete]
// @Summary Deletes like to recipe from authorized user
// @Param id path int true "Recipe id"
// @Produce json
// @Success 200
func (this *likesCtrl) del(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Likes
// @Router /recipes/{id}/like [get]
// @Summary Retrieves all liked users
// @Param id path int true "Recipe id"
// @Produce json
// @Success 200 {object} []objects.AccountDTO
func (this *likesCtrl) getByRecipe(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Likes
// @Router /accounts/like [get]
// @Summary Retrieves all liked recipes of user
// @Param login query bool false "Requested account"
// @Param me query bool false "Request by current user"
// @Produce json
// @Success 200 {object} []objects.RecipeDTO
func (this *likesCtrl) getByUser(w http.ResponseWriter, r *http.Request) {
}
