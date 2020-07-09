package main

import (
	"net/http"
	"os"

	"algogrit.com/jwt-demo/users/repository"
	"algogrit.com/jwt-demo/users/service"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var server http.Handler

const DefaultSigningKey = "5aaebec0766e57a3858ad580bb978dd38020698e8d31b03c9f22d51178644556"

func JWTMiddleware(jwtSigningKey string) func(http.Handler) http.Handler {
	options := jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSigningKey), nil
		},
		SigningMethod:       jwt.SigningMethodHS256,
		CredentialsOptional: false,
	}

	return jwtmiddleware.New(options).Handler
}

func init() {
	jwtSigningKey, ok := os.LookupEnv("JWT_SIGNING_KEY")

	if !ok {
		jwtSigningKey = DefaultSigningKey
	}

	userRepo := repository.NewUserRepo()
	userService := service.NewUserService(userRepo, jwtSigningKey)

	authMiddleware := JWTMiddleware(jwtSigningKey)

	r := mux.NewRouter()

	r.HandleFunc("/login", userService.LoginHandler).Methods("POST")
	r.Handle("/currentUser", authMiddleware(http.HandlerFunc(userService.SessionHandler))).Methods("GET")

	server = r
}

func main() {
	http.ListenAndServe(":8000", server)
}
