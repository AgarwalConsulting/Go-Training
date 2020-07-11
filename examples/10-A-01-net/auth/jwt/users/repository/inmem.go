package repository

import (
	"algogrit.com/jwt-demo/entities"
)

type UserRepo map[string]entities.User

func (ur *UserRepo) FindByUsername(username string) *entities.User {
	userMap := map[string]entities.User(*ur)
	u, ok := userMap[username]

	if !ok {
		return nil
	}

	return &u
}

func (ur *UserRepo) FindByID(userID int) *entities.User {
	userMap := map[string]entities.User(*ur)

	for _, user := range userMap {
		if user.ID == userID {
			return &user
		}
	}
	return nil
}

func NewUserRepo() *UserRepo {
	userMap := map[string]entities.User{
		"algogrit": {1, "algogrit", "drowssap", "G A", "algogrit@gmail.com"},
		"thanos":   {2, "thanos", "100%wouldbebetter", "Thanos", "thanos@hey.com"},
		"ironman":  {3, "ironman", "iloveyou3000", "RDJ", "iam@ironman.com"},
	}

	userRepo := UserRepo(userMap)
	return &userRepo
}
