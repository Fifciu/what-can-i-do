package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"

	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/fifciu/what-can-i-do/server/models"
	"github.com/gorilla/context"
)

func getProblem(w http.ResponseWriter, r *http.Request, withIdeas bool) {
	vars := mux.Vars(r)
	response := u.Status(true)
	problemSlug := strings.Trim(vars["problemSlug"], " ")

	if problemSlug == "" {
		response = u.Message(false, "Bad request")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	claims := context.Get(r, "CurrentUser")
	userId := uint(0)

	if claims != nil {
		claims := context.Get(r, "CurrentUser").(*models.Claims)
		userId = claims.ID
	}

	problem := models.GetProblem(problemSlug, withIdeas, userId)
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

func GetProblemsByQuery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	problems := models.GetProblemsByQuery(vars["searchQuery"])
	response := u.Status(true)

	if len(problems) == 0 {
		response = u.Message(false, "Nothing found")
		u.RespondWithCode(w, response, http.StatusNotFound)
		return
	} else {
		response["problems"] = problems
	}
	u.Respond(w, response)
}

func GetCertainProblem(w http.ResponseWriter, r *http.Request) {
	getProblem(w, r, false)
}

func GetCertainProblemWithIdeas(w http.ResponseWriter, r *http.Request) {
	getProblem(w, r, true)
}

func GetMostPopularProblems(w http.ResponseWriter, r *http.Request) {
	response := u.Status(true)
	response["problems"] = models.GetMostPopular()
	u.Respond(w, response)
}

func GetProblemsToReview(w http.ResponseWriter, r *http.Request) {
	response := u.Status(true)
	response["problems"] = models.GetProblemsToReview()
	u.Respond(w, response)
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

	if len(problem.Name) < 4 {
		response = u.Message(false, "Problem's name must have at least 4 characters")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}

	if len(problem.Description) < 15 {
		response = u.Message(false, "Problem's description must have at least 15 characters")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}
	claims := context.Get(r, "CurrentUser").(*models.Claims)

	if !problem.Save(claims.ID, problem.Name, problem.Description) {
		response = u.Message(false, "This problem already exists")
		u.RespondWithCode(w, response, http.StatusBadRequest)
		return
	}
	u.Respond(w, response)
}