/*
 *Author: Stefan
 *Date: 12/25/2019
 *Last changes: 01/08/2019
 *Task: Create router, build handlers, serve
 */

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/stevenstr/app"
	"github.com/stevenstr/controllers"

	"github.com/gorilla/mux"
)

//main function
func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/user/register", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.Use(app.JwtAuthentication) //attach JWT auth middleware

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
