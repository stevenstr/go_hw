/**
 *Author: Stefan
 *Date: 12/10/2019
 *Last changes: 12/11/2019 01.19
 *Task: Web server handling any request to it’s address and returning JSON like:
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//Response struct
type Response struct {
	Host       string      `json:"host"`
	UserAgent  string      `json:"user_agent"`
	RequestURI string      `json:"request_uri"`
	Headers    http.Header `json:"headers"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := &Response{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestURI: r.RequestURI,
		Headers:    r.Header,
	}

	fmt.Fprintln(w, p)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Good readability: ")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Host: ")
	json.NewEncoder(w).Encode(p.Host)
	fmt.Fprintln(w, "User_Agent")
	json.NewEncoder(w).Encode(p.UserAgent)
	fmt.Fprintln(w, "Request_URI")
	json.NewEncoder(w).Encode(p.RequestURI)
	fmt.Fprintln(w, "Headers")
	json.NewEncoder(w).Encode(p.Headers)
	fmt.Fprintln(w, "HeadersAccept: ")
	json.NewEncoder(w).Encode(p.Headers.Get("Accept"))
	fmt.Fprintln(w, "User-Agent")
	json.NewEncoder(w).Encode(p.Headers.Get("User-Agent"))
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
