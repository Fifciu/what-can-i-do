package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	u "../utils"
	"../models"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"strings"
	"time"
	"errors"
	"github.com/gorilla/context"
)

type Claims struct {
	ID uint `json:"id"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type ClaimsUser struct {
	ID uint `json:"id"`
	Fullname string `json:"fullname"`
	Email string `json:"email"`
}

var jwtSecretKey = os.Getenv("jwt_key")
var jwtKey = []byte(jwtSecretKey)

func getJwtTtl () (uint, error) {
	jwtTtlString := os.Getenv("jwt_ttl")
	jwtTtl, err := strconv.Atoi(jwtTtlString)
	if err != nil {
		return 0, errors.New("Bad value for jwt_ttl in config")
	}
	if jwtTtl == 0 {
		jwtTtl = 5
	}
	return uint(jwtTtl), nil
}

func getJwtTimeOffset () (uint, error) {
	timeOffset := os.Getenv("jwt_offset")
	timeOffsetNumber, err := strconv.Atoi(timeOffset)
	if err != nil {
		return uint(0), nil
	}
	if timeOffsetNumber == 0 {
		timeOffsetNumber = 60
	}
	return uint(timeOffsetNumber), nil
}

func generateJwt (claimsUser *ClaimsUser) (string, time.Time, error) {
	jwtTtl, err := getJwtTtl()
	if err != nil {
		return "", time.Time{}, err
	}
	expirationTime := time.Now().Add(time.Duration(jwtTtl) * time.Minute)
	claims := &Claims{
		ID: claimsUser.ID,
		Fullname: claimsUser.Fullname,
		Email: claimsUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", time.Time{}, err
	}
	return tokenString, expirationTime, nil
}

func acceptedProvider (pickedProvider string) error {
	acceptedProviders := os.Getenv("supported_oauth_providers")
	if acceptedProviders == "" {
		return errors.New("Accepted providers not configured")
	}

	providers := strings.Split(acceptedProviders, ",")
	hasRequestProvider := false
	for _, provider := range providers {
		if provider == pickedProvider {
			hasRequestProvider = true
			break
		}
	}

	if !hasRequestProvider {
		return errors.New("Not accepted provider")
	}
	return nil
}

func InitAuth(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := acceptedProvider(vars["provider"]); err != nil {
		response := u.Message(false, err.Error())
		u.RespondWithCode(w, response, http.StatusNotFound)
		return
	}

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
	vars := mux.Vars(r)

	if err := acceptedProvider(vars["provider"]); err != nil {
		response := u.Message(false, err.Error())
		u.RespondWithCode(w, response, http.StatusNotFound)
		return
	}

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

	databaseUser := &models.User{}
	newUser, err := databaseUser.CreateOrGet(user.Email(), user.Name(), vars["provider"])
	if err != nil {
		response := u.Message(false, err.Error())
		u.RespondWithCode(w, response, http.StatusInternalServerError)
		return
	}
	claimsUser := &ClaimsUser{
		ID: newUser.ID,
		Fullname: newUser.Fullname,
		Email: newUser.Email,
	}
	// Create the JWT string
	tokenString, expirationTime, err := generateJwt(claimsUser)
	if err != nil {
		response := u.Message(false, err.Error())
		u.RespondWithCode(w, response, http.StatusInternalServerError)
		return
	}
	response := u.Status(true)
	response["user"] = map[string]string{
		"email": user.Email(),
		"name": user.Name(),
	}
	response["token"] = tokenString
	response["expires_at"] = expirationTime

	u.RespondWithCode(w, response, http.StatusOK)
}

func RefreshToken (w http.ResponseWriter, r *http.Request) {
	timeOffsetNumber, err := getJwtTimeOffset()
	if err != nil {
		response := u.Message(false, err.Error())
		u.RespondWithCode(w, response, http.StatusNotFound)
		return
	}

	claims := context.Get(r, "CurrentUser").(*Claims)

	// Expired offset
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) < time.Duration(timeOffsetNumber)*time.Minute*-1 {
		u.RespondWithCode(w, map[string]interface{}{"status": false}, http.StatusBadRequest)
		return
	}

	claimsUser := &ClaimsUser{
		ID: claims.ID,
		Fullname: claims.Fullname,
		Email: claims.Email,
	}
	// Create the JWT string
	tokenString, expirationTime, err := generateJwt(claimsUser)
	if err != nil {
		response := u.Message(false, err.Error())
		u.RespondWithCode(w, response, http.StatusInternalServerError)
		return
	}
	response := u.Status(true)
	response["token"] = tokenString
	response["expires_at"] = expirationTime

	u.RespondWithCode(w, response, http.StatusOK)
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	claims := context.Get(r, "CurrentUser").(*Claims)
	response := u.Status(true)
	response["user"] = map[string]interface{}{
		"fullname": claims.Fullname,
		"email": claims.Email,
	}
	u.RespondWithCode(w, response, http.StatusOK)
	return
}