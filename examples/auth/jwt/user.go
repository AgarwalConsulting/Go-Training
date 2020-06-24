package main

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}

func (u *User) ComparePassword(creds *User) bool {
	if u == nil {
		return false
	}

	return u.Password == creds.Password
}
