package main

import (
	"fmt"
	"net/http"
	"os"

	controllers "github.com/fifciu/what-can-i-do/server/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/",
		controllers.HelloWorld).Methods("GET")

	router.HandleFunc("/problems",
		controllers.GetProblemByQuery).Methods("GET").Queries("searchQuery", "{searchQuery}")

	router.HandleFunc("/problems",
		controllers.GetProblems).Methods("GET")

	router.HandleFunc("/problems/{problemId:[0-9]+}",
		controllers.GetCertainProblem).Methods("GET")

	router.HandleFunc("/problems/{problemId:[0-9]+}/ideas",
		controllers.GetCertainProblemWithIdeas).Methods("GET")

	router.HandleFunc("/problems",
		controllers.AddProblem).Methods("POST")

	router.HandleFunc("/ideas",
		controllers.AddIdea).Methods("POST")

	port := os.Getenv("PORT")
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