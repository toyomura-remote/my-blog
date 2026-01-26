package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64
	Name     string
	Did      string
	Email    string
	Password string
	Post     []*Post
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func NewUser(email, hashedPassword string) (*User, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	user := User{
		Did:      uuid.String(),
		Email:    email,
		Password: hashedPassword,
	}

	return &user, nil
}
