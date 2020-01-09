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

*/

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/FogCreek/mini"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

//fatal function
func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//params function
func params() string {

	//C:\Users\admin\go\src\github.com\stevenstr  u.HomeDir
	cfg, err := mini.LoadConfiguration("C:/Users/admin/go/src/github.com/stevenstr/.tweeeet")
	fatal(err)

	info := fmt.Sprintf("host=%s port=%s dbname=%s "+
		"sslmode=%s user=%s password=%s ",
		cfg.String("host", "127.0.0.1"),
		cfg.String("port", "5432"),
		cfg.String("dbname", "postgres"),
		cfg.String("sslmode", "disable"),
		cfg.String("user", "postgres"),
		cfg.String("pass", "1"),
	)
	return info
}

var db *sql.DB

//main function
func main() {

	var err error
	db, err = sql.Open("postgres", params())
	fatal(err)

	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " +
		`tweets("id" SERIAL PRIMARY KEY,` +
		`"message" text, "account_id" text)`)

	fatal(err)

	router := httprouter.New()
	router.GET("/api/v1/tweets", getTweets)
	router.GET("/api/v1/tweets/:id", getTweet)
	router.POST("/api/v1/tweets", addTweet)

	http.ListenAndServe(":8080", router)
}
