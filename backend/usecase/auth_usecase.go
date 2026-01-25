package usecase

import (
	"context"
	"my-blog-backend/domain"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	SignUp(ctx context.Context, email string, password string) error
	Login(ctx context.Context, email string, password string) (string, error)
	GetUserByEmail(ctx context.Context, token string) (*domain.User, error)
}

type authUseCase struct {
	authRepo domain.AuthRepository
}

func NewAuthUseCase(authRepo domain.AuthRepository) AuthUseCase {
	return &authUseCase{authRepo: authRepo}
}

func (u *authUseCase) SignUp(ctx context.Context, email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	user := domain.User{
		Did:      uuid.String(),
		Email:    email,
		Password: string(hashedPassword),
	}
	return u.authRepo.CreateUser(ctx, user)
}

func (u *authUseCase) Login(ctx context.Context, email string, password string) (string, error) {
	return "", nil
}
func (u *authUseCase) GetUserByEmail(ctx context.Context, token string) (*domain.User, error) {
	return nil, nil
}
