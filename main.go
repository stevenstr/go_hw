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
	Host       string `json:"host"`
	UserAgent  string `json:"user_agent"`
	RequestURI string `json:"request_uri"`
	HeaderU    string `json:"headers"`
	HeaderA    string `json:"Accept"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	p := &Response{
		Host:       r.Host,
		UserAgent:  r.UserAgent(),
		RequestURI: r.URL.String(),
		HeaderU:    r.Header.Get("User-Agent"),
		HeaderA:    r.Header.Get("Accept"),
	}

	fmt.Fprintln(w, *p)

	fmt.Fprintln(w)
	fmt.Fprintln(w, "Good readability: ")
	fmt.Fprintln(w)

	json.NewEncoder(w).Encode(p.Host)
	json.NewEncoder(w).Encode(p.UserAgent)
	json.NewEncoder(w).Encode(p.RequestURI)
	json.NewEncoder(w).Encode(p.HeaderU)
	json.NewEncoder(w).Encode(p.HeaderA)
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
