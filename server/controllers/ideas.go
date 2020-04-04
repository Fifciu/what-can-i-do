package controllers

import (
	"encoding/json"
	"net/http"

	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/fifciu/what-can-i-do/server/models"
)

func AddIdea(w http.ResponseWriter, r *http.Request) {
	idea := &models.Idea{}
	err := json.NewDecoder(r.Body).Decode(idea)
	response := u.Status(true)

	if err != nil {
		response = u.Message(false, "Could not add an idea")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	if idea.ProblemID < 1 {
		response = u.Message(false, "Bad problem ID")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	if !models.ProblemExists(idea.ProblemID) {
		response = u.Message(false, "Problem does not exist")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	if len(idea.Description) < 15 {
		response = u.Message(false, "Problem's description must have at least 15 characters")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	if !idea.Save(idea.Description, idea.Price, idea.ProblemID) {
		response = u.Message(false, "Problem with this description has already exist")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}
	u.Respond(w, response)
}
