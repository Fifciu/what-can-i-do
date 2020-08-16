package controllers

import (
	"net/http"

	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/fifciu/what-can-i-do/server/models"
)

func GetIdeasToReview(w http.ResponseWriter, r *http.Request) {
	response := u.Status(true)
	response["ideas"] = models.GetIdeasToReview()
	u.Respond(w, response)
}