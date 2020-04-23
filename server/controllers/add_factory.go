package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/context"
	u "github.com/fifciu/what-can-i-do/server/utils"
)

func AddRecordFactory(entity models.DatabaseType) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(entity)
		response := u.Status(true)

		// No input
		if err != nil {
			response = u.Message(false, "Data not provided")
			u.RespondWithCode(w, response, http.StatusBadRequest)
			return
		}

		claims := context.Get(r, "CurrentUser").(*Claims)
		entity.UserID = claims.ID

		// Bad input
		if err := entity.Validate(); err != nil {
			response = u.Message(false, err.Error())
			u.RespondWithCode(w, response, http.StatusBadRequest)
			return
		}

		if err := entity.Save(); err != nil {
			response = u.Message(false, err.Error())
			u.RespondWithCode(w, response, http.StatusBadRequest)
			return
		}
		u.Respond(w, response)
	})
}