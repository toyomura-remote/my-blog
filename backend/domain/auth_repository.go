package domain

import (
	"context"
	dto "my-blog-backend/DTO"
)

type AuthRepository interface {
	CreateUser(ctx context.Context, user *User) (*dto.UserResponse, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	ExistsByEmail(ctx context.Context, email string) (bool, error)
}
