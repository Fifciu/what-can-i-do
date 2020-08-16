package middlewares

import (
	"net/http"
	"github.com/fifciu/what-can-i-do/server/models"
	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/gorilla/context"
)


func Moderator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := context.Get(r, "CurrentUser").(*models.Claims)
		if (claims.Flags & (1<<0)) != 1 {
			u.RespondWithCode(w, map[string]interface{}{"status": false}, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
