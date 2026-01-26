package main

import (
	"database/sql"
	"my-blog-backend/infra/db"
	"my-blog-backend/lib/wire"
	middlewares "my-blog-backend/middleware"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Postgresドライバー
)

func setUpRouter(conn *sql.DB) *gin.Engine {

	authUseCase := wire.InitAuthUseCase(conn)
	authController := wire.InitAuthController(conn)
	postController := wire.InitPostController(conn)

	r := gin.Default()
	r.RedirectTrailingSlash = true // 末尾スラッシュを自動リダイレクト
	authRouter := r.Group("/auth")
	// postRouter := r.Group("/post")
	postRouterWithAuth := r.Group("/post", middlewares.AuthMiddleware(authUseCase))

	authRouter.POST("/signup", authController.SignUp)
	authRouter.POST("/login", authController.Login)

	postRouterWithAuth.POST("", postController.CreatePost)
	postRouterWithAuth.GET("", postController.GetPostsByUserID)

	return r
}

func main() {
	db.Initialize()
	conn := db.SetupDB()
	defer conn.Close()
	r := setUpRouter(conn)

	println("==== Ser! ====")
	r.Run()
}
