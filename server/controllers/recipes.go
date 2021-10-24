package controllers

import (
	"api/recipes/models"
	"api/recipes/repository"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type recipe struct {
	db *gorm.DB
}

func InitRecipes(r *mux.Router, db *gorm.DB) {
	ctrl := &recipe{db}
	r.HandleFunc("/recipes", ctrl.getAllRecipes).Methods("GET")
}

// @Tags Recipes
// @Router /recipes [get]
// @Summary Retrieves all recipes
// @Produce json
// @Success 200 {object} []objects.Recipes
func (this *recipe) getAllRecipes(w http.ResponseWriter, r *http.Request) {
	model := models.NewRecipe(repository.NewRecipesRep(this.db))
	data := model.GetAll()
	jsonSuccess(w, data)
}