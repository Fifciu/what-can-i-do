package utils

import (
	"os"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func InitOauthProviders (pwaBaseUrl string) {
	gomniauth.SetSecurityKey("SOME_AUTH_KEY")
	googleClientId := os.Getenv("google_client_id")
	googleClientSecret := os.Getenv("google_client_secret")
	gomniauth.WithProviders(
		google.New(googleClientId, googleClientSecret, pwaBaseUrl + "auth/google"),
	)
}
