package main

type UserRepo map[string]User

func (ur *UserRepo) FindByUsername(username string) *User {
	userMap := map[string]User(*ur)
	u, ok := userMap[username]

	if !ok {
		return nil
	}

	return &u
}

func (ur *UserRepo) FindByID(userID int) *User {
	userMap := map[string]User(*ur)

	for _, user := range userMap {
		if user.ID == userID {
			return &user
		}
	}
	return nil
}

func NewUserRepo() *UserRepo {
	userMap := map[string]User{
		"algogrit": {1, "algogrit", "drowssap", "G A", "algogrit@gmail.com"},
		"thanos":   {2, "thanos", "100%wouldbebetter", "Thanos", "thanos@hey.com"},
		"ironman":  {3, "ironman", "iloveyou3000", "RDJ", "iam@ironman.com"},
	}

	userRepo := UserRepo(userMap)
	return &userRepo
}
