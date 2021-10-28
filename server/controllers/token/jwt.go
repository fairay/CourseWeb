package auth

import (
	"api/recipes/controllers"
	"api/recipes/utils"
	"net/http"
	"time"
	"regexp"

	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	Login	string
	jwt.StandardClaims
}

func NewToken(login string) (*Token, time.Time) {
	expTime := time.Now().Add(30 * time.Minute)
	return &Token{
			Login: login,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expTime.Unix(),
			},
		},
		expTime
}

func (tk *Token) ToString() string {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenStr, err := token.SignedString([]byte(utils.Config.TokenPassword))
	if err != nil {
		tokenStr = ""
	}

	return tokenStr
}

func FillCookie(w http.ResponseWriter, login string) {
	tk, expTime := NewToken(login)
	tokenStr := tk.ToString()

	if len(tokenStr) == 0 {
		controllers.BadRequest(w, "Error in formatting token to string")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenStr,
		Expires: expTime,
		Path:    "/",

		HttpOnly: true,
	})
}

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{`^/accounts/.*`, `^/swagger/.*`}//,  "/accounts/register", "/accounts/login"} 
		requestPath := r.URL.Path

		for _, value := range notAuth {
			isMatch, _ := regexp.MatchString(value, requestPath)
			if isMatch {
				next.ServeHTTP(w, r)
				return
			}
		}

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				controllers.TokenIsMissed(w)
				return
			}

			controllers.BadRequest(w, "Can not extract token")
			return
		}

		tokenStr := c.Value
		tk := &Token{}

		token, err := jwt.ParseWithClaims(tokenStr, tk, func(token *jwt.Token) (interface{}, error) {
			return utils.Config.TokenPassword, nil
		})

		if err != nil || !token.Valid { 
			controllers.AccessDenied(w)
			return
		}

		if time.Now().Unix()-tk.ExpiresAt > utils.Config.TokenLiveTime {
			controllers.TokenExpired(w)
			return
		}

		FillCookie(w, tk.Login)

		next.ServeHTTP(w, r)
	}) 
	} 