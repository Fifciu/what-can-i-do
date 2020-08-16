/*
	Loool
 */
package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/fifciu/what-can-i-do/server/models"
	"strconv"
	"github.com/gorilla/context"
)

type payload struct {
	Accepted bool `json:"accepted"`
	Message string `json:"message"`
}

// ResolveFactory return a HTTP Handler for resolving new idea/problem by moderator
func ResolveFactory(entityType string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// add review
		// update status if needed
		vars := mux.Vars(r)
		entityId, err := strconv.Atoi(vars["problem_id"])
		response := u.Status(true)
		if err != nil || entityId < 1 {
			response = u.Message(false, "Bad request")
			u.RespondWithCode(w, response, http.StatusBadRequest)
			return
		}
		payload := &payload{}
		err = json.NewDecoder(r.Body).Decode(payload)
		if err != nil {
			response = u.Message(false, "Bad request")
			u.RespondWithCode(w, response, http.StatusBadRequest)
			return
		}
		if payload.Accepted {
			if entityType == "problem" {
				entity := &models.Problem{}
				entity.ID = uint(entityId)
				err = entity.Resolve()
				if err != nil {
					response = u.Message(false, err.Error())
					u.RespondWithCode(w, response, http.StatusInternalServerError)
					return
				}
				u.Respond(w, response)
				return

			} else if entityType == "idea" {
				entity := &models.Idea{}
				entity.ID = uint(entityId)
				err = entity.Resolve()
				if err != nil {
					response = u.Message(false, err.Error())
					u.RespondWithCode(w, response, http.StatusInternalServerError)
					return
				}
				u.Respond(w, response)
				return

			} else {
				response = u.Message(false, "Internal server error")
				u.RespondWithCode(w, response, http.StatusInternalServerError)
				return
			}
		} else {
			if len(payload.Message) < 5 {
				response = u.Message(false, "Feedback message should have at least 5 chars")
				u.RespondWithCode(w, response, http.StatusBadRequest)
				return
			}
			claims := context.Get(r, "CurrentUser").(*models.Claims)
			if entityType == "problem" || entityType == "idea" {
				if entityType == "problem" {
					review := &models.ProblemReview{}
					err = review.Save(uint(entityId), claims.ID, payload.Message)
				} else if entityType == "idea" {
					review := &models.IdeaReview{}
					err = review.Save(uint(entityId), claims.ID, payload.Message)
				}
				if err != nil {
					response = u.Message(false, "Internal server error")
					u.RespondWithCode(w, response, http.StatusInternalServerError)
					return
				}
			} else {
				response = u.Message(false, "Internal server error")
				u.RespondWithCode(w, response, http.StatusInternalServerError)
				return
			}
		}
		u.Respond(w, response)
	})
}