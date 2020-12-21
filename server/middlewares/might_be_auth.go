package middlewares

import (
	"net/http"
	"strings"
	"github.com/fifciu/what-can-i-do/server/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)
// AAAAAA
func MightBeAuthUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		splitedHeader := strings.Split(authHeader, " ")
		token := strings.Trim(splitedHeader[1], " ")

		claims := &models.Claims{}
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})


		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				next.ServeHTTP(w, r)
				return
			}
			if strings.HasPrefix(err.Error(), "token is expired by") && r.URL.Path == "/auth/refresh" {
				// Allow to try to refresh
				context.Set(r, "CurrentUser", claims)
				next.ServeHTTP(w, r)
				return
			}
			next.ServeHTTP(w, r)
			return
		}

		if !tkn.Valid {
			next.ServeHTTP(w, r)
			return
		}

		// How to share claims between? Context?
		context.Set(r, "CurrentUser", claims)
		next.ServeHTTP(w, r)
	})
}
