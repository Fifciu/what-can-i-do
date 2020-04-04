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
		controllers.GetProblems).Methods("GET")

	router.HandleFunc("/problems/{problemId}",
		controllers.GetCertainProblem).Methods("GET")

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