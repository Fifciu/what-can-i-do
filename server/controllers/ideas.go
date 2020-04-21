package controllers

import (
	"encoding/json"
	"net/http"

	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/fifciu/what-can-i-do/server/models"
	"github.com/gorilla/context"
)

func GetMineIdeas(w http.ResponseWriter, r *http.Request) {
	claims := context.Get(r, "CurrentUser").(*Claims)
	ideas := models.GetUserIdeas(claims.ID, []models.IdeasMapper{models.MapperAddProblemsName})
	response := u.Status(true)
	response["ideas"] = ideas
	u.RespondWithCode(w, response, http.StatusOK)
}

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

	if len(idea.ActionDescription) < 15 {
		response = u.Message(false, "Idea's action description must have at least 15 characters")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	if len(idea.ResultsDescription) < 15 {
		response = u.Message(false, "Idea's results description must have at least 15 characters")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	if idea.MoneyPrice < 0 {
		response = u.Message(false, "Idea's price be bigger or equal $0")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	claims := context.Get(r, "CurrentUser").(*Claims)

	if !idea.Save(claims.ID, idea.ProblemID, idea.ActionDescription, idea.ResultsDescription, idea.MoneyPrice, idea.TimePrice) {
		response = u.Message(false, "Problem with this description has already exist")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}
	u.Respond(w, response)
}
