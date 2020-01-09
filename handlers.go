/*
*Author: Stefan
*Date: 12/18/2019
*Last changes: 12/18/2019 13.55
*Task: functions for json
*/

package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//Tweet struct
type Tweet struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	AID     string `json:"account_id"`
}

//для поиска по аккаунт айди
func getID(w http.ResponseWriter, ps httprouter.Params) (int, bool) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(400)
		return 0, false
	}
	return id, true
}

//getTweets functin
func getTweets(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var str string
	if len(r.URL.RawQuery) > 0 {
		str = r.URL.Query().Get("message")
		if str == "" {
			w.WriteHeader(400)
			return
		}
	}
	recs, err := read(str)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(recs); err != nil {
		w.WriteHeader(500)
	}
}

//getTweet function
func getTweet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, ok := getID(w, ps)
	if !ok {
		return
	}
	rec, err := readOne(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			return
		}
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err = json.NewEncoder(w).Encode(rec); err != nil {
		w.WriteHeader(500)
	}
}

//addTweet function
func addTweet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var rec Tweet
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil || rec.Message == "" || rec.AID == "" {
		w.WriteHeader(400)
		return
	}
	if _, err := insert(rec.Message, rec.AID); err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
}
