package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/fifciu/what-can-i-do/server/models"
	"github.com/gorilla/context"
)

func GetIdeaReviews(w http.ResponseWriter, r *http.Request) {
	response := u.Status(true)

	vars := mux.Vars(r)
	entityId, err := strconv.Atoi(vars["idea_id"])
	if err != nil {
		response["status"] = false
		response["message"] = "Bad idea id"
		u.Respond(w, response)
		return
	}
	claims := context.Get(r, "CurrentUser").(*models.Claims)

	response["reviews"] = models.GetIdeaReviews(uint(entityId), claims.ID)
	u.Respond(w, response)
}

func GetProblemReviews(w http.ResponseWriter, r *http.Request) {
	response := u.Status(true)

	vars := mux.Vars(r)
	entityId, err := strconv.Atoi(vars["problem_id"])
	if err != nil {
		response["status"] = false
		response["message"] = "Bad idea id"
		u.Respond(w, response)
		return
	}
	claims := context.Get(r, "CurrentUser").(*models.Claims)

	response["reviews"] = models.GetProblemReviews(uint(entityId), claims.ID)
	u.Respond(w, response)
}