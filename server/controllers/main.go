package controllers

import (
	"net/http"

	u "../utils"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	response := u.Message(true, "Hello!")
	u.Respond(w, response)
}