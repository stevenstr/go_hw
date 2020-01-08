/*
 *Author: Stefan
 *Date: 12/25/2019
 *Last changes: 01/08/2019
 *Task: Create account and Authenticate
 */

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/stevenstr/models"
	u "github.com/stevenstr/utils"
)

//CreateAccount var
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Create account
	u.Respond(w, resp)
}

//Authenticate var
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
