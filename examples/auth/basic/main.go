package main

import (
	"net/http"
)

const (
	username = "admin"
	password = "123456"
)

func BasicAuth(username, password, realm string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
		iusername, ipassword, _ := req.BasicAuth()

		if username != iusername || password != ipassword {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		next(w, req)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	http.HandleFunc("/", BasicAuth(username, password, "Please enter your username and password for this site", handleIndex))

	http.ListenAndServe(":9000", nil)
}
