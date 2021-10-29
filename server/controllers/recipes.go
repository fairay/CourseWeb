package controllers

import (
	"api/recipes/controllers/responses"
	auth "api/recipes/controllers/token"
	"api/recipes/models"
	"api/recipes/objects"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type recipe struct {
	model *models.RecipeM
}

func InitRecipes(r *mux.Router, model *models.RecipeM) {
	ctrl := &recipe{model}
	r.HandleFunc("/recipes", ctrl.getAllRecipes).Methods("GET")
	r.HandleFunc("/recipes/add", ctrl.addRecipe).Methods("POST")
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

// TODO: Check me out (add or new?) // add or post
// @Tags Recipes
// @Router /recipes/add [post]
// @Param recipe body objects.RecipeDTO false "Recipe data"
// @Summary Creation a new recipe
// @Produce json
// @Success 200 {object} objects.RecipeDTO
func (this *recipe) addRecipe(w http.ResponseWriter, r *http.Request) {
	rcpDTO := &objects.RecipeDTO{}
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

	/*TODO: стоит скорее всего передавать объект + вопрос про то, что в 
	* дто нет категорий, а их нужно всё-таки 
	*/
	err = this.model.AddRecipe(rcpDTO.ToModel())
}
