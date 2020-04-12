package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	u "github.com/fifciu/what-can-i-do/server/utils"
	"github.com/fifciu/what-can-i-do/server/models"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"strings"
	"time"
)

type Claims struct {
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var jwtSecretKey = os.Getenv("jwt_key")
var jwtKey = []byte(jwtSecretKey)

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

	databaseUser := &models.User{}
	databaseUser.Save(user.Email(), user.Name(), vars["provider"])

	jwtTtlString := os.Getenv("jwt_ttl")
	jwtTtl, err := strconv.Atoi(jwtTtlString)
	if err != nil {
		response := u.Message(false, "Could not read JWT TTL")
		u.RespondWithCode(w, response, http.StatusInternalServerError)
		return
	}
	if jwtTtl == 0 {
		jwtTtl = 5
	}
	expirationTime := time.Now().Add(time.Duration(jwtTtl) * time.Minute)
	claims := &Claims{
		Fullname: user.Name(),
		Email: user.Email(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response["token"] = tokenString
	response["expires_in"] = expirationTime.Sub(time.Now()).Milliseconds()

	u.RespondWithCode(w, response, http.StatusOK)
}