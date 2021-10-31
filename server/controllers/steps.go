package controllers

import (
	"api/recipes/models"
	"net/http"

	"github.com/gorilla/mux"
)

type stepCtrl struct {
	model *models.StepM
}

func InitSteps(r *mux.Router, model *models.StepM) {
	ctrl := &stepCtrl{model}
	r.HandleFunc("/recipes/{id}/steps", ctrl.getAll).Methods("GET")
	r.HandleFunc("/recipes/{id}/steps", ctrl.post).Methods("POST")

	r.HandleFunc("/recipes/{id}/steps/{step}", ctrl.get).Methods("GET")
	r.HandleFunc("/recipes/{id}/steps/{step}", ctrl.patch).Methods("PATCH")
	r.HandleFunc("/recipes/{id}/steps/{step}", ctrl.del).Methods("DELETE")
}

// TODO:
// @Tags Steps
// @Router /recipes/{id}/steps [get]
// @Summary Retrieves all steps
// @Param id path int true "Recipe's id"
// @Produce json
// @Success 200 {object} []objects.StepDTO
func (this *stepCtrl) getAll(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Steps
// @Router /recipes/{id}/steps [post]
// @Summary Posts step
// @Param id path int true "Recipe's id"
// @Param step body objects.StepDTO true "Step"
// @Produce json
// @Success 201 {object} objects.StepDTO
func (this *stepCtrl) post(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Steps
// @Router /recipes/{id}/steps/{step} [get]
// @Summary Retrieves step
// @Param id path int true "Recipe's id"
// @Param step path int true "Step's number"
// @Produce json
// @Success 200 {object} objects.StepDTO
func (this *stepCtrl) get(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Steps
// @Router /recipes/{id}/steps/{step} [patch]
// @Summary Updates step
// @Param id path int true "Recipe's id"
// @Param step path int true "Step's number"
// @Produce json
// @Success 200 {object} objects.StepDTO
func (this *stepCtrl) patch(w http.ResponseWriter, r *http.Request) {
}

// TODO:
// @Tags Steps
// @Router /recipes/{id}/steps/{step} [delete]
// @Summary Removes step
// @Param id path int true "Recipe's id"
// @Param step path int true "Step's number"
// @Produce json
// @Success 200
func (this *stepCtrl) del(w http.ResponseWriter, r *http.Request) {
}
