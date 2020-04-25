package controllers

import (
	"net/http"
	"github.com/gorilla/context"
	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/fifciu/what-can-i-do/server/models"
)

func GetMineFactory(entity models.UserCreatedEntity) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := context.Get(r, "CurrentUser").(*models.Claims)
		entities := entity.GetByUserId(claims.ID)
		response := u.Status(true)
		response[entity.PluralName()] = entities
		u.Respond(w, response)
	})
}