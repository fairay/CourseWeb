package controllers

import (
	"encoding/json"
	"net/http"
)

func jsonSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(data)
}

func TokenIsMissed(w http.ResponseWriter) {
	msg := "Missing auth token"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(msg)
}

func AccessDenied(w http.ResponseWriter) {
	msg := "jwt-token is not valid"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(msg)
}

func TokenExpired(w http.ResponseWriter) {
	msg := "jwt-token expired"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(msg)
}

func AuthenticationFailed(w http.ResponseWriter) {
	msg := "Authentication failed"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(msg)
}

func BadRequest(w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(msg)
}
