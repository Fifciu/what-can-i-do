package main

import (
	"fmt"
	"net/http"
	"os"

	controllers "github.com/fifciu/what-can-i-do/server/controllers"
	//middlewares "github.com/fifciu/what-can-i-do/server/middlewares"

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

	gomniauth.SetSecurityKey("SOME_AUTH_KEY")
	gomniauth.WithProviders(
		google.New("key", "secret", "callback"),
	)

	router := mux.NewRouter()

	router.HandleFunc("/",
		controllers.HelloWorld).Methods("GET")

	router.HandleFunc("/problems",
		controllers.GetProblemsByQuery).Methods("GET").Queries("searchQuery", "{searchQuery}")

	router.HandleFunc("/problems",
		controllers.GetProblems).Methods("GET")

	router.HandleFunc("/problems/{problemId:[0-9]+}",
		controllers.GetCertainProblem).Methods("GET")

	router.Handle("/problems/{problemId:[0-9]+}/ideas",
			http.HandlerFunc(controllers.GetCertainProblemWithIdeas)).Methods("GET")

	router.HandleFunc("/problems",
		controllers.AddProblem).Methods("POST")

	router.HandleFunc("/ideas",
		controllers.AddIdea).Methods("POST")

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