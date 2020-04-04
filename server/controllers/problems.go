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
	code := 200

	if err != nil {
		response = u.Message(false, "Problem does not exist")
		code = http.StatusBadRequest
	}

	if !problem.Save(problem.Title, problem.Description) {
		response = u.Message(false, "Problem with this title has already exists")
		code = http.StatusBadRequest
	}
	u.RespondWithCode(w, response, code)
}