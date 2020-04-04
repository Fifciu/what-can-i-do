package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/fifciu/what-can-i-do/server/models"
)

func GetProblems(w http.ResponseWriter, r *http.Request) {
	response := u.Status(true)
	response["problems"] = models.GetAllProblems()
	u.Respond(w, response)
}

func GetCertainProblem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	response := u.Status(true)
	problemId, err := strconv.Atoi(vars["problemId"])
	if problemId < 1 || err != nil {
		response = u.Message(false, "Bad request")
		w.WriteHeader(http.StatusBadRequest)
	}

	problem := models.GetProblem(problemId, false)
	if problem.ID == 0 {
		response = u.Message(false, "Problem does not exist")
		w.WriteHeader(http.StatusNotFound)
	} else {
		response["problem"] = problem
		response["sdsd"] = problem.ID
	}

	u.Respond(w, response)
}