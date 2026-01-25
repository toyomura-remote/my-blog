package infra

import (
	"context"
	"database/sql"
	"my-blog-backend/domain"
	"my-blog-backend/infra/model"

	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) domain.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) CreateUser(ctx context.Context, user domain.User) error {
	insertUser := model.User{
		Did:      user.Did,
		Name:     user.Name,
		Password: user.Password,
	}

	if err := insertUser.Insert(ctx, r.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	modelUser, err := model.Users(
		qm.Where("email = ?", email),
	).One(ctx, r.db)

	if err != nil {
		return nil, err
	}

	domainUser := &domain.User{
		Name: modelUser.Name,
		// Did:  modelUser.,
	}

	return domainUser, nil
}
