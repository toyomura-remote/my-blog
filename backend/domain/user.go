package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64   `json:"-"`
	Name     string  `json:"name"`
	Did      string  `json:"did"`
	Email    string  `json:"-"`
	Password string  `json:"-"`
	Post     []*Post `json:"post"`
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func NewUser(name, email, hashedPassword string) (*User, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	user := User{
		Did:      uuid.String(),
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	return &user, nil
}
