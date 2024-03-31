package models

import (
	"errors"
	"slices"
	"time"

	"github.com/AMCodeBytes/go-chat/utilities"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `form:"username" json:"username"`
	Password  string    `form:"password" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

var users = []User{}

func (user User) Login() (string, error) {
	idx := slices.IndexFunc(users, func(u User) bool { return u.Username == user.Username })

	if idx == -1 {
		return "", errors.New("no user exists")
	}

	u := &users[idx]

	match := utilities.AuthenticatePassword(user.Password, u.Password)

	if !match {
		return "", errors.New("invalid credentials")
	}

	return u.ID, nil
}

func (user User) Create() {
	users = append(users, user)
}
