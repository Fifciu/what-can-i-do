package main

import (
	"fmt"
	"net/http"
	"os"

	controllers "github.com/fifciu/what-can-i-do/server/controllers"
	models "github.com/fifciu/what-can-i-do/server/models"
	middlewares "github.com/fifciu/what-can-i-do/server/middlewares"
	u "github.com/fifciu/what-can-i-do/server/utils"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// TODO First query after launch does not work. No connection with db?

	pwaProtocol := os.Getenv("pwa_protocol")
	pwaHost := os.Getenv("pwa_host")
	pwaPort := os.Getenv("pwa_port")
	pwaBaseUrl := pwaProtocol + "://" + pwaHost
	if pwaPort != "" {
		pwaBaseUrl = pwaBaseUrl + ":" + pwaPort
	}
	pwaBaseUrl = pwaBaseUrl + "/"

	u.InitOauthProviders(pwaBaseUrl)

	router := mux.NewRouter()

	// Home's view - Search query
	router.HandleFunc("/problems",
		controllers.GetProblemsByQuery).Methods("GET").Queries("searchQuery", "{searchQuery}")

	//router.HandleFunc("/problems",
	//	controllers.GetProblems).Methods("GET")

	// Account's view
	router.Handle("/problems/mine",
		middlewares.AuthUser(controllers.GetMineFactory(&models.Problem{}))).Methods("POST")

	//router.HandleFunc("/problems/{problemSlug}",
	//	controllers.GetCertainProblem).Methods("GET")

	// Problem's view
	router.Handle("/problems/{problemSlug}/ideas",
			middlewares.MightBeAuthUser(http.HandlerFunc(controllers.GetCertainProblemWithIdeas))).Methods("GET")

	// Home's view - Add new problem
	router.Handle("/problems",
		middlewares.AuthUser(http.HandlerFunc(controllers.AddProblem))).Methods("POST")

	// Problem's view - Add new idea
	router.Handle("/ideas",
		middlewares.AuthUser(controllers.AddRecordFactory(&models.Idea{}))).Methods("POST")

	// Problem's view - Add vote up/down
	router.Handle("/vote",
		middlewares.AuthUser(controllers.AddRecordFactory(&models.Vote{}))).Methods("POST")

	// Account's view - My ideas
	router.Handle("/ideas/mine",
		middlewares.AuthUser(controllers.GetMineFactory(&models.Idea{}))).Methods("POST")

	// Sign in/up's view - Init auth and get redirect link
	router.HandleFunc("/auth/init/{provider}",
		controllers.InitAuth).Methods("POST")

	// After redirection's view -
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