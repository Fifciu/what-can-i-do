package controllers

import (
	"net/http"
	"github.com/gorilla/context"
	u "../utils"
	"../models"
)

type UserCreatedEntity interface {
	GetByUserId(userId uint) []*UserCreatedEntity
	PluralName() string
}

func GetMineFactory(entity models.UserCreatedEntity) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := context.Get(r, "CurrentUser").(*Claims)
		entities := entity.GetByUserId(claims.ID)
		response[entity.PluralName()] = entities
		u.Respond(w, response)
	})
}