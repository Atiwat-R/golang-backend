package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID int
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterUserPayload struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
}








