package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/fifciu/what-can-i-do/server/models"
)

func getProblem(w http.ResponseWriter, r *http.Request, withIdeas bool) {
	vars := mux.Vars(r)
	response := u.Status(true)
	problemId, err := strconv.Atoi(vars["problemId"])

	if problemId < 1 || err != nil {
		response = u.Message(false, "Bad request")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	problem := models.GetProblem(problemId, withIdeas)
	if problem.ID == 0 {
		response = u.Message(false, "Problem does not exist")
		u.RespondWithCode(w, response, http.StatusNotFound)
		return
	} else {
		response["problem"] = problem
	}

	u.Respond(w, response)
}

func GetProblems(w http.ResponseWriter, r *http.Request) {
	response := u.Status(true)
	response["problems"] = models.GetAllProblems()
	u.Respond(w, response)
}

func GetProblemByQuery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	problem := models.GetProblemByQuery(vars["searchQuery"])
	response := u.Status(true)

	if problem.ID == 0 {
		response = u.Message(false, "Problem does not exist")
		u.RespondWithCode(w, response, http.StatusNotFound)
		return
	} else {
		response["problem"] = problem
	}
	u.Respond(w, response)
}

func GetCertainProblem(w http.ResponseWriter, r *http.Request) {
	getProblem(w, r, false)
}

func GetCertainProblemWithIdeas(w http.ResponseWriter, r *http.Request) {
	getProblem(w, r, true)
}

func AddProblem(w http.ResponseWriter, r *http.Request) {
	problem := &models.Problem{}
	err := json.NewDecoder(r.Body).Decode(problem)
	response := u.Status(true)

	if err != nil {
		response = u.Message(false, "Couldn't not add problem")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	if len(problem.Title) < 4 {
		response = u.Message(false, "Problem's title must have at least 4 characters")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	if len(problem.Description) < 15 {
		response = u.Message(false, "Problem's description must have at least 15 characters")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	if !problem.Save(problem.Title, problem.Description) {
		response = u.Message(false, "Problem with this title has already exists")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}
	u.Respond(w, response)
}