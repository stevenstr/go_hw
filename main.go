/**
 *Author: Stefan
 *Date: 12/12/2019
 *Last changes: 12/11/2019 11.05
 *Task: Web server handling any request to it’s address and returning JSON like:
 */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//Response struct
type Response struct {
	Host       string   `json:"host"`
	UserAgent  string   `json:"user_agent"`
	RequestURI string   `json:"request_uri"`
	Headers    struct { //is correctly accoding the tasks
		Accept    string
		UserAgent string
	} `json:"headers"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := Response{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestURI: r.RequestURI,
		Headers: struct {
			Accept    string
			UserAgent string
		}{
			r.Header.Get("Accept"),
			r.Header.Get("User-agent"),
		},
	}

	//pack for Write(interface)
	pp, err := json.Marshal(&p)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(pp)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", jsonHandler)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at: 8080")
	server.ListenAndServe()
}
