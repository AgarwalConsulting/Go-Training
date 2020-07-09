package service

import (
	"encoding/json"
	"net/http"

	"algogrit.com/jwt-demo/entities"
	"algogrit.com/jwt-demo/users/repository"
	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

type UserService struct {
	Repo          *repository.UserRepo
	JWTSigningKey string
}

func (us UserService) getToken(u *entities.User) map[string]string {
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
	var creds entities.User

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

func (us UserService) SessionHandler(w http.ResponseWriter, req *http.Request) {
	token := req.Context().Value("user").(*jwt.Token)
	claims := map[string]interface{}(token.Claims.(jwt.MapClaims))
	userID := int(claims["userID"].(float64))
	user := us.Repo.FindByID(userID)

	log.Info("User: ", user)

	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func NewUserService(repo *repository.UserRepo, jwtSigningKey string) *UserService {
	return &UserService{repo, jwtSigningKey}
}
