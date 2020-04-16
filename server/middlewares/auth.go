package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/fifciu/what-can-i-do/server/controllers"
	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = os.Getenv("jwt_key")
var jwtKey = []byte(jwtSecretKey)

func AuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
				u.RespondWithCode(w, map[string]interface{}{"status": false}, http.StatusUnauthorized)
				return
		}

		splitedHeader := strings.Split(authHeader, " ")
		token := strings.Trim(splitedHeader[1], " ")

		claims := &controllers.Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				u.RespondWithCode(w, map[string]interface{}{"status": false}, http.StatusUnauthorized)
				return
			}
			u.RespondWithCode(w, map[string]interface{}{"status": false}, http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			u.RespondWithCode(w, map[string]interface{}{"status": false}, http.StatusUnauthorized)
			return
		}

		// How to share claims between? Context?
		w.Write([]byte(fmt.Sprintf("Welcome %d!", claims.ID)))

	})
}
