package auth

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	USERNAME = "algogrit"
	PASSWORD = "1234567890"
)

// BasicAuthMiddleware performs basic auth before forwarding request
func BasicAuthMiddleware(h http.Handler) http.Handler {
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		inputUsername, inputPassword, ok := r.BasicAuth()

		if !ok {
			realm := "Please enter your username and password for this site!"
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
		}

		if !(inputUsername == USERNAME && inputPassword == PASSWORD) {
			log.Error("Unsucessful authentication: ", inputUsername, inputPassword)
			// w.WriteHeader(http.StatusUnauthorized)
			// w.Write([]byte("Unauthorized"))

			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		log.Info("authenticated successfully!")
		h.ServeHTTP(w, r) // healthzHandler, employeeIndexHandler
	}

	return http.HandlerFunc(handlerFunc)
}

// BasicAuthMiddleware performs basic auth before forwarding request
// func BasicAuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	basicAuthMiddleware(http.HandlerFunc(next)).ServeHTTP(w, r)
// }
