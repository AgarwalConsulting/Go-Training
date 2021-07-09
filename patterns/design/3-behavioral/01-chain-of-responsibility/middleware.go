package main

import "net/http"

func authenticationMiddlerware(h http.Handler) http.Handler {
	authFn := func(w http.ResponseWriter, r *http.Request) {
		authenticated := true
		if authenticated {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}

	return http.HandlerFunc(authFn)
}

func loggingMiddlerware(h http.Handler) http.Handler {
	return h
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	// ...
	HelloWorldHandler := http.HandlerFunc(HelloWorld)

	loggedHelloHandler := loggingMiddlerware(HelloWorldHandler)

	authLoggedHelloHandler := authenticationMiddlerware(loggedHelloHandler)

	http.ListenAndServe(":8080", authLoggedHelloHandler)
}
