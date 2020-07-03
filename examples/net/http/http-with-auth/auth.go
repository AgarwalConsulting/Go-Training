package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	USERNAME = "algogrit"
	PASSWORD = "1234567890"
)

func basicAuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	inputUsername, inputPassword, ok := r.BasicAuth()

	if !ok {
		realm := "Please enter your username and password for this site!"
		w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
	}

	if !(inputUsername == USERNAME && inputPassword == PASSWORD) {
		log.Error("Unsucessful authentication: ", inputUsername, inputPassword)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	log.Info("authenticated successfully!")
	next(w, r) // healthzHandler, employeeIndexHandler
}
