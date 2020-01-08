/*
*Author: Stefan
 *Date: 12/18/2019
 *Last changes: 12/18/2019 13.55
 *Task:

			Level low
			POST /tweets   - create a tweet
				Payload:
					message - some tweet message
					account_id - any number to distinguish accounts
				Result:
					id - message primary key
					message - tweet message
					account_id - number to distinguish accounts

			GET    /tweets?account_id=“your account id”    -  return all tweets expect yours order by timestamp newest first

				Result:
					tweets -  all tweets expect yours order by timestamp newest first


				RUN go get github.com/gorilla/mux
*/
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Tweet struct
type Tweet struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	AccountID string `json:"accountid"`
}

var tweets []Tweet

//createTweet struct
func createTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tweet Tweet
	_ = json.NewDecoder(r.Body).Decode(&tweet)
	tweet.ID = strconv.Itoa(rand.Intn(1000000))
	tweets = append(tweets, tweet)
	json.NewEncoder(w).Encode(tweet)

}

//getTweet struct
func getTweet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range tweets {
		if item.AccountID == params["accountid"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Tweet{})
}

//getTweets struct
func getTweets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tweets)
}

//main func
func main() {
	r := mux.NewRouter()

	tweets = append(tweets, Tweet{ID: "1", Message: "ferger", AccountID: "Stylus"})
	tweets = append(tweets, Tweet{ID: "2", Message: "ererge", AccountID: "Johnny123"})
	//some handlers
	r.HandleFunc("/tweets/{accountid}", getTweet).Methods("GET")
	r.HandleFunc("/tweets", createTweet).Methods("POST")
	r.HandleFunc("/tweetss", getTweets).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
