package main

import (
	"fmt"
	"net/http"
	"os"

	controllers "github.com/fifciu/what-can-i-do/server/controllers"
	models "github.com/fifciu/what-can-i-do/server/models"
	middlewares "github.com/fifciu/what-can-i-do/server/middlewares"
	u "github.com/fifciu/what-can-i-do/server/utils"

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

	// Home's view - Hot
	router.HandleFunc("/problems/hot",
		controllers.GetMostPopularProblems).Methods("GET")

	// Problems to review
	router.Handle("/problems/review-request",
		middlewares.AuthUser(middlewares.Moderator(http.HandlerFunc(controllers.GetProblemsToReview)))).Methods("GET")

	// Ideas to review
	router.Handle("/ideas/review-request",
		middlewares.AuthUser(middlewares.Moderator(http.HandlerFunc(controllers.GetIdeasToReview)))).Methods("GET")

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

	//// Author - Get idea reviews
	// TODO: Prevention mechanism for multiple reviews
	router.Handle("/ideas/{idea_id:[0-9]+}/reviews",
		middlewares.AuthUser(http.HandlerFunc(controllers.GetIdeaReviews))).Methods("GET")
	//// Author - Get problem reviews
	// TODO: Prevention mechanism for multiple reviews
	router.Handle("/problems/{problem_id:[0-9]+}/reviews",
		middlewares.AuthUser(http.HandlerFunc(controllers.GetIdeaReviews))).Methods("GET")

	//// Moderator - Review idea
	// TODO: Prevention mechanism for multiple reviews
	router.Handle("/ideas/{idea_id:[0-9]+}/review",
		middlewares.AuthUser(middlewares.Moderator(controllers.ResolveFactory("idea")))).Methods("POST")
	//// Moderator - Review problem
	// TODO: Prevention mechanism for multiple reviews
	router.Handle("/problems/{problem_id:[0-9]+}/review",
		middlewares.AuthUser(middlewares.Moderator(controllers.ResolveFactory("problem")))).Methods("POST")

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

	// TODO Prepare nginx configuration for same domain api & pwa
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}

