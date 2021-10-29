package auth

import (
	"net/http"
	jwt "github.com/dgrijalva/jwt-go"

	"api/recipes/utils"
)

// TODO: Check me out
func LoginFromCookie(r *http.Request) (string, error) {
	c, err := r.Cookie("token")
	if err != nil {
		return "", err
	}

	tokenStr := c.Value
	tk := &Token{}

	_, err = jwt.ParseWithClaims(tokenStr, tk, func(token *jwt.Token) (interface{}, error) {
		return utils.Config.TokenPassword, nil
	})

	return tk.Login, err
}