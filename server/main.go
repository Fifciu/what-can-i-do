package main

import (
	"fmt"
	"net/http"
	"os"

	controllers "github.com/fifciu/what-can-i-do/server/controllers"
	middlewares "github.com/fifciu/what-can-i-do/server/middlewares"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
)

func main() {

	//User:
	//	Id
	//	Fullname
	//	Email
	//	Password
	//	CreatedAt
	//	<smth from fb/google/linkedin?>
	//
	//Vote:
	//	Id
	//	IdeaID
	//	IsPlus

	pwaProtocol := os.Getenv("pwa_protocol")
	pwaHost := os.Getenv("pwa_host")
	pwaPort := os.Getenv("pwa_port")
	pwaBaseUrl := pwaProtocol + "://" + pwaHost
	if pwaPort != "" {
		pwaBaseUrl = pwaBaseUrl + ":" + pwaPort
	}
	pwaBaseUrl = pwaBaseUrl + "/"

	gomniauth.SetSecurityKey("SOME_AUTH_KEY")
	googleClientId := os.Getenv("google_client_id")
	googleClientSecret := os.Getenv("google_client_secret")
	gomniauth.WithProviders(
		google.New(googleClientId, googleClientSecret, pwaBaseUrl + "auth/google"),
	)

	router := mux.NewRouter()

	router.HandleFunc("/",
		controllers.HelloWorld).Methods("GET")

	router.Handle("/protected",
		middlewares.AuthUser(http.HandlerFunc(controllers.HelloWorld))).Methods("GET")

	router.HandleFunc("/problems",
		controllers.GetProblemsByQuery).Methods("GET").Queries("searchQuery", "{searchQuery}")

	router.HandleFunc("/problems",
		controllers.GetProblems).Methods("GET")

	router.HandleFunc("/problems/{problemSlug}",
		controllers.GetCertainProblem).Methods("GET")

	router.Handle("/problems/{problemSlug}/ideas",
			http.HandlerFunc(controllers.GetCertainProblemWithIdeas)).Methods("GET")

	router.Handle("/problems",
		middlewares.AuthUser(http.HandlerFunc(controllers.AddProblem))).Methods("POST")

	router.Handle("/ideas",
		middlewares.AuthUser(http.HandlerFunc(controllers.AddIdea))).Methods("POST")

	router.HandleFunc("/auth/init/{provider}",
		controllers.InitAuth).Methods("POST")

	router.HandleFunc("/auth/complete/{provider}",
		controllers.CompleteAuth).Methods("POST")

	router.Handle("/auth/refresh",
		middlewares.AuthUser(http.HandlerFunc(controllers.RefreshToken))).Methods("POST")

	router.Handle("/me",
		middlewares.AuthUser(http.HandlerFunc(controllers.GetMe))).Methods("POST")

	port := os.Getenv("api_port")
	if port == "" {
		port = "8090"
	}

	fmt.Println("It works on http://localhost:" + port + "/ address")

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"POST", "GET"})

	err := http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router))
	if err != nil {
		fmt.Print(err)
	}

}