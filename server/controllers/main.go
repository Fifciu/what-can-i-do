package controllers

import (
	"net/http"

	u "github.com/fifciu/what-can-i-do/server/utils"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	response := u.Message(true, "Hello!")
	u.Respond(w, response)
}