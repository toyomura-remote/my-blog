package usecase

import (
	"context"
	"fmt"
	"my-blog-backend/domain"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, email, password string) error
	Login(ctx context.Context, email string, password string) (*string, error)
	GetUserFromToken(ctx context.Context, tokenString string) (*domain.User, error)
}

type authUseCase struct {
	authRepo domain.AuthRepository
}

func NewAuthUseCase(authRepo domain.AuthRepository) AuthUseCase {
	return &authUseCase{authRepo: authRepo}
}

func (u *authUseCase) SignUp(ctx context.Context, email, password string) error {

	hashedPassword, err := domain.HashPassword(password)
	if err != nil {
		return err
	}

	user, err := domain.NewUser(email, hashedPassword)

	if err := u.authRepo.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *authUseCase) Login(ctx context.Context, email string, password string) (*string, error) {
	foundUser, err := u.authRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password)); err != nil {
		return nil, err
	}

	token, err := CreateToken(uint(foundUser.ID), foundUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil

}

// TODO:infraに切り出して呼ぶ際はdomainに作成するインターフェースを参照する
func CreateToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(([]byte(os.Getenv("SECRET_KEY"))))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (s *authUseCase) GetUserFromToken(ctx context.Context, tokenString string) (*domain.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	var user *domain.User
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, jwt.ErrTokenExpired
		}

		user, err = s.authRepo.GetUserByEmail(ctx, claims["email"].(string))
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}
