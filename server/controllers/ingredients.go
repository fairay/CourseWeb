package controllers

import (
	"api/recipes/models"
	"net/http"

	"github.com/gorilla/mux"
)

type ingredientCtrl struct {
	model *models.RecipeM
}

func InitIngredients(r *mux.Router, model *models.RecipeM) {
	ctrl := &ingredientCtrl{model}

	r.HandleFunc("/recipes/{id}/ingredients", ctrl.getByRecipe).Methods("GET")
	r.HandleFunc("/recipes/{id}/ingredients", ctrl.putToRecipe).Methods("PUT")
	r.HandleFunc("/recipes/{id}/ingredients", ctrl.postToRecipe).Methods("POST")
	r.HandleFunc("/recipes/{id}/ingredients", ctrl.delFromRecipe).Methods("DELETE")
}


// TODO:
// @Tags Ingredients
// @Router /recipes/{id}/ingredients [get]
// @Summary Retrieves all ingredients
// @Param id path int true "Recipe's id"
// @Produce json
// @Success 200 {object} []objects.IngredientDTO
func (this *ingredientCtrl) getByRecipe(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Ingredients
// @Router /recipes/{id}/ingredients [post]
// @Summary Posts all ingredients
// @Param id path int true "Recipe's id"
// @Param recipes body []objects.IngredientDTO true "Ingredients"
// @Produce json
// @Success 201 {object} []objects.IngredientDTO
func (this *ingredientCtrl) postToRecipe(w http.ResponseWriter, r *http.Request) {
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
