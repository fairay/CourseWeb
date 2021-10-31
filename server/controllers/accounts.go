package controllers

import (
	"api/recipes/controllers/responses"
	auth "api/recipes/controllers/token"
	"api/recipes/models"
	"api/recipes/objects"
	"time"

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
	r.HandleFunc("/accounts/logout", ctrl.LogOut).Methods("POST")
	r.HandleFunc("/accounts", ctrl.addAccount).Methods("POST")
	r.HandleFunc("/accounts/{login}", ctrl.getAccount).Methods("GET")
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
		responses.BadRequest(w, "Invalid request")
		return
	}

	acc, err := this.model.LogIn(accDTO.Login, accDTO.Password)
	if err != nil {
		responses.AuthenticationFailed(w)
		return
	}

	auth.FillCookie(w, acc.Login)
}

// @Tags Accounts
// @Router /accounts/logout [post]
// @Summary Logging out
// @Produce json
// @Success 200
func (this *account) LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
		Path:    "/",

		HttpOnly: true,
	})

	responses.TextSuccess(w, "Logout was successful")
}

// TODO:
// @Tags Accounts
// @Router /accounts [post]
// @Param account body objects.AccountDTO true "Account data"
// @Summary Creation a new account
// @Produce json
// @Success 201 {object} objects.AccountDTO
func (this *account) addAccount(w http.ResponseWriter, r *http.Request) {
	// _ := new(objects.AccountDTO)
}

// @Tags Accounts
// @Router /accounts/{login} [get]
// @Summary Retrieves account
// @Param login path string true "Account login"
// @Produce json
// @Success 200 {object} objects.AccountDTO
func (this *account) getAccount(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	login := urlParams["login"]

	data, _ := this.model.Find(login)
	responses.JsonSuccess(w, data.ToDTO())
}
