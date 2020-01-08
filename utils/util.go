/*
 *Author: Stefan
 *Date: 12/25/2019
 *Last changes: Just utils for sending status message
 */

package utils

import (
	"encoding/json"
	"net/http"
)

//Message function
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

//Respond function
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
