package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var server http.Handler

const DefaultSigningKey = "5aaebec0766e57a3858ad580bb978dd38020698e8d31b03c9f22d51178644556"

func init() {
	jwtSigningKey, ok := os.LookupEnv("JWT_SIGNING_KEY")

	if !ok {
		jwtSigningKey = DefaultSigningKey
	}

	userRepo := NewUserRepo()
	userService := NewUserService(userRepo, jwtSigningKey)

	r := mux.NewRouter()

	r.HandleFunc("/login", userService.LoginHandler).Methods("POST")

	server = r
}

func main() {
	http.ListenAndServe(":8000", server)
}
