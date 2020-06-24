package main

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	Repo          *UserRepo
	JWTSigningKey string
}

func (us UserService) getToken(u *User) map[string]string {
	token := jwt.New(jwt.SigningMethodHS256)

	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)

	/* Set token claims */
	claims["name"] = u.Name
	claims["userID"] = u.ID

	// log.Info(us.JWTSigningKey)
	/* Sign the token with our secret */
	tokenString, err := token.SignedString([]byte(us.JWTSigningKey))

	if err != nil {
		log.Fatal(err)
	}

	tokenMap := map[string]string{"token": tokenString}

	return tokenMap
}

func (us UserService) LoginHandler(w http.ResponseWriter, req *http.Request) {
	var creds User

	json.NewDecoder(req.Body).Decode(&creds)

	user := us.Repo.FindByUsername(creds.Username)

	if !user.ComparePassword(&creds) {
		http.Error(w, "cannot login - incorrect username or password", http.StatusUnauthorized)
		return
	}

	// Generate a jwt token
	token := us.getToken(user)

	json.NewEncoder(w).Encode(token)
}

func NewUserService(repo *UserRepo, jwtSigningKey string) *UserService {
	return &UserService{repo, jwtSigningKey}
}
