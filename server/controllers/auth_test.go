package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	u "github.com/fifciu/what-can-i-do/server/utils"
)

type InitAuthResponse struct {
	Status bool
	RedirectUrl string
	Message string
}

func TestGenerateJWT (t *testing.T) {
	// Arrange
	claimsUser := &ClaimsUser{
		ID:       1,
		Fullname: "John Doe",
		Email:    "john@doe.com",
	}

	// Act
	_, expirationTime, _ := generateJwt(claimsUser)

	// Assert
	jwtTtl, err := getJwtTtl()
	if err != nil {
		t.Errorf("getJwtTtl did not return jwtttl")
	}
	sub := expirationTime.Sub(time.Now()).Minutes()
	if uint(math.Round(sub)) != jwtTtl {
		t.Errorf("Bad expiry time for the token")
	}
}

func TestAcceptedProvider (t *testing.T) {
	// Arrange
	providersResults := map[string]bool{
		"google": true,
		"tuptack": false,
	}

	for provider, expect := range providersResults {
		// Act
		err := 	acceptedProvider(provider)
		// Assert
		if expect && err != nil {
			t.Errorf("Bad results, expected %s to be proper option", provider)
		} else if !expect && err == nil {
			t.Errorf("Bad results, expected %s to be invalid option", provider)
		}
	}
}

func TestInitAuth (t *testing.T) {
	req, err := http.NewRequest("POST", "/auth/init/bad-provider", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	req2, err := http.NewRequest("POST", "/auth/init/google", nil)
	req2 = mux.SetURLVars(req2, map[string]string{"provider": "google"})
	if err != nil {
		t.Fatal(err.Error())
	}
	u.InitOauthProviders("http://localhost:8090")

	rr := httptest.NewRecorder()
	rr2 := httptest.NewRecorder()

	handler := http.HandlerFunc(InitAuth)

	handler.ServeHTTP(rr, req)
	handler.ServeHTTP(rr2, req2)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Bad response code, got %v wanted %v", status, http.StatusNotFound)
	}

	responseEnt := &InitAuthResponse{}
	err = json.NewDecoder(rr2.Body).Decode(responseEnt)
	if err != nil {
		t.Errorf(err.Error())
	}

	if !responseEnt.Status || len(responseEnt.RedirectUrl) < 10 {
		t.Errorf("Did not return auth link when it should")
	}
}