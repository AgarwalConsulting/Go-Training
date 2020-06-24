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

func NewUserRepo() *UserRepo {
	userMap := map[string]User{
		"algogrit": {1, "algogrit", "2020isjumanji", "G A", "algogrit@gmail.com"},
		"thanos":   {2, "thanos", "100%wouldbebetter", "Thanos", "thanos@hey.com"},
		"ironman":  {3, "ironman", "iloveyou3000", "RDJ", "iam@ironman.com"},
	}

	userRepo := UserRepo(userMap)
	return &userRepo
}
