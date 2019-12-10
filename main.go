/**
 *Author: Stefan
 *Date: 12/10/2019
 *Last changes: 12/10/2019 17.00
 *Task:
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
	Header     http.Header `json:"headers"`
}

//Our handler
func handler(w http.ResponseWriter, r *http.Request) {

	//1 method
	//json.NewEncoder(w).Encode(Response{r.Host, r.UserAgent(), r.RequestURI, r.Header})

	//2 method
	p := &Response{r.Host, r.UserAgent(), r.RequestURI, r.Header}
	json.Marshal(p)
	fmt.Fprintln(w, p)
	/* 3 method
	fmt.Fprintln(w, "URL:", r.URL.String())
	fmt.Fprintln(w, "host:", r.Host)
	fmt.Fprintln(w, "user_agent:", r.UserAgent())
	fmt.Fprintln(w, "Accept:", r.Header.Get("Accept"))
	fmt.Fprintln(w, "User-Agent:", r.Header.Get("User-Agent"))
	*/
}

func main() {
	//обработчики
	//http.HandleFunc("/", handler)
	//напишет в консоли где запущен сервак

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at: 8080")
	//запускает сервер на порту 8080
	server.ListenAndServe()
}
