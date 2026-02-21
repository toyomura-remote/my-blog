package infra

import (
	"context"
	"database/sql"
	dto "my-blog-backend/DTO"
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

// TODO: ExistsByEmailを作成する

func (r *authRepository) CreateUser(ctx context.Context, user *domain.User) (*dto.UserResponse, error) {
	insertUser := model.User{
		Did:      user.Did,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if err := insertUser.Insert(ctx, r.db, boil.Infer()); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:    uint(insertUser.ID),
		Name:  insertUser.Name,
		Email: insertUser.Email,
	}, nil

}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	modelUser, err := model.Users(
		qm.Where("email = ?", email),
	).One(ctx, r.db)

	if err != nil {
		return nil, err
	}

	domainUser := &domain.User{
		ID:       modelUser.ID,
		Name:     modelUser.Name,
		Did:      modelUser.Did,
		Email:    modelUser.Email,
		Password: modelUser.Password,
	}

	return domainUser, nil
}

func (r *authRepository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	return model.Users(
		qm.Where("email = ?", email),
	).Exists(ctx, r.db)
}
