/*
 *Author: Stefan
 *Date: 12/11/2019
 *Last changes: 12/11/2019 16.20
 *Task:
• Web server, responds to “/” (404 on any other request)
• Support POST and GET methods
• On GET request returns HTML page, containing FORM with name and address fields + submit button.
Also there is a placeholder to display token on this page.
• On POST request reads name+address from body and creates a token as “name:address”.
Then saves this token to Cookies
• Page reloads, token value displayed on page
**/

package main

import (
	"fmt"
	"net/http"
	"time"
)

var formNameAddrTmpl = []byte(`
	<html>
	<head>
		<meta charset="UTF-8" />
	</head>
	<body>
	<div>Known token: <span id="known-token"></span></div>
	<div>
		<form method="POST" action="/">
			<label>Name</label><input name="name" type="text" value="" />
			<label>Address</label><input name="address" type="text" value="" />
			<input type="submit" value="submit" />
		</form>
	</div>
	</body>
	<script type="text/javascript">
		const elem = document.querySelector("#known-token");
		let token = document.cookie.split(';').filter((item) => item.trim().startsWith('token='))[0];
		elem.innerHTML = token.replace('token=', '')
	</script>
	</html>
`)

func handleMainPage(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Write(formNameAddrTmpl)
		return

	case http.MethodPost:
		name := r.FormValue("name")
		address := r.FormValue("address")
		t := time.Now().Add(10 * time.Hour)
		cookie := http.Cookie{Name: "token", Value: name + ":" + address, Expires: t}
		http.SetCookie(w, &cookie)
		w.Write(formNameAddrTmpl)
		http.Redirect(w, r, "/", http.StatusFound)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found"))
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleMainPage)

	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at: 8080")
	server.ListenAndServe()
}
