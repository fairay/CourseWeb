package controllers

import (
	"api/recipes/models"
	"api/recipes/objects"
	"auth"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type account struct {
	model *models.AccountM
}

func InitAccount(r *mux.Router, model *models.AccountM) {
	ctrl := &account{model}
	r.HandleFunc("/accounts/login", ctrl.LogIn).Methods("POST")
}

// @Tags Accounts
// @Router /accounts/login [post]
// @Param account body objects.AccountDTO false "Authentication data"
// @Summary User's authorization
// @Produce json
// @Success 200 {object} objects.AccountDTO
func (this *account) LogIn(w http.ResponseWriter, r *http.Request) {
	accDTO := &objects.AccountDTO{}

	err := json.NewDecoder(r.Body).Decode(accDTO)
	if err != nil {
		BadRequest(w, "Invalid request")
		return
	}

	acc, err := this.model.LogIn(accDTO.Login, accDTO.Password)
	if err != nil {
		AuthenticationFailed(w)
		return
	}

	auth.FillCookie(w, acc.Login)
}
