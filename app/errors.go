/*
 *Author: Stefan
 *Date: 12/25/2019
 *Last changes: 01/08/2019
 *Task: JWT
 */

package app

import (
	u "github.com/stevenstr/utils"
	"net/http"
)

//NotFoundHandler var
var NotFoundHandler = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		u.Respond(w, u.Message(false, "This resources was not found"))
		next.ServeHTTP(w, r)
	})
}
