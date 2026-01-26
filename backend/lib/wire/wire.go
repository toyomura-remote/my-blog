//go:build wireinject
// +build wireinject

package wire

import (
	"database/sql"
	"my-blog-backend/controller"
	"my-blog-backend/infra"
	"my-blog-backend/usecase"

	"github.com/google/wire"
)

func InitAuthController(db *sql.DB) controller.AuthController {
	wire.Build(
		infra.NewAuthRepository,
		usecase.NewAuthUseCase,
		controller.NewAuthController,
	)
	return nil
}

func InitAuthUseCase(conn *sql.DB) usecase.AuthUseCase {
	wire.Build(
		infra.NewAuthRepository,
		usecase.NewAuthUseCase,
	)
	return nil
}

func InitPostController(db *sql.DB) controller.PostController {
	wire.Build(
		infra.NewPostRepository,
		usecase.NewPostUseCase,
		controller.NewPostController,
	)
	return nil
}
