package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/context"
	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/fifciu/what-can-i-do/server/models"
)

func AddRecordFactory(entity models.DatabaseType) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		copyEntity := entity.GetNewInstance()
		err := json.NewDecoder(r.Body).Decode(copyEntity)
		response := u.Status(true)

		// No input
		if err != nil {
			response = u.Message(false, err.Error())
			u.RespondWithCode(w, response, http.StatusBadRequest)
			return
		}

		claims := context.Get(r, "CurrentUser").(*Claims)
		copyEntity.SetUserId(claims.ID)

		// Bad input
		if err := copyEntity.Validate(); err != nil {
			response = u.Message(false, err.Error())
			u.RespondWithCode(w, response, http.StatusBadRequest)
			return
		}

		if err := copyEntity.Save(); err != nil {
			response = u.Message(false, err.Error())
			u.RespondWithCode(w, response, http.StatusBadRequest)
			return
		}
		response["entity"] = copyEntity
		u.Respond(w, response)
	})
}