package controllers

import (
	"api/recipes/models"
	"api/recipes/repository"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type category struct {
	db *gorm.DB
}

func InitCategories(r *mux.Router, db *gorm.DB) {
	ctrl := &category{db}
	r.HandleFunc("/categories", ctrl.getAllCategories).Methods("GET")
	r.HandleFunc("/categories/find/{req}", ctrl.getCertainCategories).Methods("GET")
}

// @Tags Categories
// @Router /categories [get]
// @Summary Retrieves all categories
// @Produce json
// @Success 200 {object} []objects.Categories
func (this *category) getAllCategories(w http.ResponseWriter, r *http.Request) {
	model := models.NewCategory(repository.NewCategotiesRep(this.db))
	data := model.GetAll()
	jsonSuccess(w, data)
}

// @Tags Categories
// @Router /categories/find/{req} [get]
// @Summary Retrieves categories with certain item
// @Param req path string false "Search query"
// @Produce json
// @Success 200 {object} []objects.Categories
func (this *category) getCertainCategories(w http.ResponseWriter, r *http.Request) {

	model := models.NewCategory(repository.NewCategotiesRep(this.db))
	temp := mux.Vars(r)
	data := model.Find(temp["req"])
	jsonSuccess(w, data)
}
