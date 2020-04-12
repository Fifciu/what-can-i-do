package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	u "github.com/fifciu/what-can-i-do/server/utils"
	"os"
	"strings"
)

func acceptedProvider (w http.ResponseWriter, r *http.Request) bool {
	vars := mux.Vars(r)
	acceptedProviders := os.Getenv("supported_oauth_providers")
	if acceptedProviders == "" {
		response := u.Message(false, "Accepted providers not configured")
		u.RespondWithCode(w, response, http.StatusInternalServerError)
		return false
	}

	providers := strings.Split(acceptedProviders, ",")
	hasRequestProvider := false
	for _, provider := range providers {
		if provider == vars["provider"] {
			hasRequestProvider = true
			break
		}
	}

	if !hasRequestProvider {
		response := u.Message(false, "Not accepted provider")
		u.RespondWithCode(w, response, http.StatusInternalServerError)
		return false
	}
	return true
}

func InitAuth(w http.ResponseWriter, r *http.Request) {

	if !acceptedProvider(w, r) {
		return
	}

	vars := mux.Vars(r)

	provider, err := gomniauth.Provider(vars["provider"])
	if err != nil {
		response := u.Message(false, "Could not init authenticate process")
		u.RespondWithCode(w, response, http.StatusNotFound)
		return
	}
	authUrl, err := provider.GetBeginAuthURL(nil, nil)
	if err != nil {
		response := u.Message(false, "Could not create redirect URL")
		u.RespondWithCode(w, response, http.StatusNotFound)
		return
	}
	response := u.Status(true)
	response["redirectUrl"] = authUrl
	u.RespondWithCode(w, response, http.StatusOK)
}

func CompleteAuth(w http.ResponseWriter, r *http.Request) {
	if !acceptedProvider(w, r) {
		return
	}
	vars := mux.Vars(r)

	provider, err := gomniauth.Provider(vars["provider"])
	if err != nil {
		response := u.Message(false, "Could not select this provider")
		u.RespondWithCode(w, response, http.StatusInternalServerError)
		return
	}
	//queryParams, err := objx.FromURLQuery(r.URL.RawQuery)
	//if err != nil {
	//	response := u.Message(false, "Malformed URL")
	//	u.RespondWithCode(w, response, http.StatusBadRequest)
	//	return
	//}
	tmp := map[string]interface{}{"code": r.URL.Query().Get("code")}
	queryParams := objx.Map(tmp)
	creds, err := provider.CompleteAuth(queryParams)
	if err != nil {
		fmt.Println((err))
		response := u.Message(false, "Could not complete auth")
		u.RespondWithCode(w, response, http.StatusInternalServerError)
		return
	}
	user, err := provider.GetUser(creds)
	if err != nil {
		response := u.Message(false, "Could not find user for this code")
		u.RespondWithCode(w, response, http.StatusInternalServerError)
		return
	}
	response := u.Status(true)
	response["user"] = map[string]string{
		"email": user.Email(),
		"name": user.Name(),
		"nickname": user.Nickname(),
	}
	u.RespondWithCode(w, response, http.StatusOK)
}