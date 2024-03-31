package models

import "time"

type User struct {
	ID        string    `json:"id"`
	Username  string    `form:"username" json:"username"`
	Password  string    `form:"password" json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

var users = []User{}

func (user User) Create() {
	users = append(users, user)
}
