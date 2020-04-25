package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	u "../utils"
	"../models"
	"github.com/gorilla/context"
	"github.com/dgrijalva/jwt-go"
)

type InitAuthResponse struct {
	Status bool
	RedirectUrl string
	Message string
}

type RefreshTokenResponse struct {
	Status bool
	Token string
	ExpiresAt time.Time
}

type GetMeResponse struct {
	Status bool
	User models.User
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

func TestRefreshToken (t *testing.T) {
	req, err := http.NewRequest("POST", "/refresh", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	req2, err := http.NewRequest("POST", "/refresh", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	u.InitOauthProviders("http://localhost:8090")
	jwtOffset, err := getJwtTimeOffset()
	if err != nil {
		t.Fatal(err.Error())
	}
	expiredTime := time.Now().Add(time.Duration(jwtOffset) * time.Minute * -2)
	context.Set(req, "CurrentUser", &Claims{
		ID: 2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredTime.Unix(),
		},
	})
	context.Set(req2, "CurrentUser", &Claims{
		ID: 2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix(),
		},
	})

	rr := httptest.NewRecorder()
	rr2 := httptest.NewRecorder()

	handler := http.HandlerFunc(RefreshToken)

	handler.ServeHTTP(rr, req)
	handler.ServeHTTP(rr2, req2)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Bad response code, got %v wanted %v", status, http.StatusBadRequest)
	}

	responseEnt := &RefreshTokenResponse{}
	err = json.NewDecoder(rr2.Body).Decode(responseEnt)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !responseEnt.Status || len(responseEnt.Token) < 10 {
		t.Errorf("Did not return proper token and date")
	}
}

func TestGetMe (t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	expectedFullname := "Fif Jot"
	expectedEmail := "fif@gmail.com"
	context.Set(req, "CurrentUser", &Claims{
		ID: 2,
		Fullname: expectedFullname,
		Email: expectedEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix(),
		},
	})

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMe)

	// Act
	handler.ServeHTTP(rr, req)

	// Assert
	responseEnt := &GetMeResponse{}
	err = json.NewDecoder(rr.Body).Decode(responseEnt)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !responseEnt.Status || responseEnt.User.Email != expectedEmail || responseEnt.User.Fullname != expectedFullname {
		t.Errorf("Did not return proper token and date")
	}
}