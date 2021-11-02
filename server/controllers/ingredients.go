package controllers

import (
	"api/recipes/controllers/responses"
	"api/recipes/errors"
	"api/recipes/models"
	"api/recipes/objects"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ingredientCtrl struct {
	model *models.IngredientM
}

func InitIngredients(r *mux.Router, model *models.IngredientM) {
	ctrl := &ingredientCtrl{model}

	r.HandleFunc("/recipes/{id}/ingredients", ctrl.getByRecipe).Methods("GET")
	r.HandleFunc("/recipes/{id}/ingredients", ctrl.putToRecipe).Methods("PUT")
	r.HandleFunc("/recipes/{id}/ingredients", ctrl.postToRecipe).Methods("POST")
	r.HandleFunc("/recipes/{id}/ingredients", ctrl.delFromRecipe).Methods("DELETE")
}

// TODO:
// @Tags Ingredients
// @Router /recipes/{id}/ingredients [get]
// @Summary Retrieves all recipe's ingredients
// @Param id path int true "Recipe's id"
// @Produce json
// @Success 200 {object} []objects.IngredientDTO
// @Failure 400 Invalid value
func (this *ingredientCtrl) getByRecipe(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]

	id_rcp, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Wrong recipe's id")
		return
	}

	data, err := this.model.GetByRecipe(id_rcp)

	switch err {
	case nil:
		responses.JsonSuccess(w, data)
	case errors.UnknownRecipe:
		responses.RecordNotFound(w, "recipe")
	default:
		responses.BadRequest(w, "Error in getting ingredients")
	}
}

// TODO:
// @Tags Ingredients
// @Router /recipes/{id}/ingredients [post]
// @Summary Posts all recipe's ingredients
// @Param id path int true "Recipe's id"
// @Param recipes body []objects.IngredientDTO true "Ingredients"
// @Produce json
// @Success 201 Successful opeartion
// @Failure 400 Invalid value
func (this *ingredientCtrl) postToRecipe(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	strId := urlParams["id"]

	id_rcp, err := strconv.Atoi(strId)
	if err != nil {
		responses.BadRequest(w, "Wrong recipe's id")
		return
	}

	ingDTO := new([]objects.IngredientDTO)
	err = json.NewDecoder(r.Body).Decode(ingDTO)
	if err != nil {
		responses.BadRequest(w, "Invalid request")
		return
	}

	err = this.model.PostToRecipe(id_rcp, objects.ToArrModelI(*ingDTO))
	switch err {
	case nil:
		responses.SuccessCreation(w, "Posting ingredient/s was successful")
	case errors.UnknownRecipe:
		responses.BadRequest(w, "Wrong recipe's id")
	case errors.DBAdditionError:
		responses.BadRequest(w, "Error in addition a new ingredient")
	default:
		responses.BadRequest(w, "Invalid request")
	}
}

// TODO:
// @Tags Ingredients
// @Router /recipes/{id}/ingredients [put]
// @Summary Adds ingredient
// @Param id path int true "Recipe's id"
// @Param recipe body objects.IngredientDTO true "Ingredient"
// @Produce json
// @Success 200
func (this *ingredientCtrl) putToRecipe(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Ingredients
// @Router /recipes/{id}/ingredients [delete]
// @Summary Removes ingredient
// @Param id path int true "Recipe's id"
// @Param recipe body objects.IngredientDTO true "Ingredient"
// @Produce json
// @Success 200
func (this *ingredientCtrl) delFromRecipe(w http.ResponseWriter, r *http.Request) {
}
