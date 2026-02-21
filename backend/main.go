package main

import (
	"database/sql"
	"my-blog-backend/infra/db"
	"my-blog-backend/lib/wire"
	"my-blog-backend/lib/wire/config"
	middlewares "my-blog-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // Postgresドライバー
)

func setUpRouter(conn *sql.DB) *gin.Engine {

	authUseCase := wire.InitAuthUseCase(conn)
	authController := wire.InitAuthController(conn)
	postController := wire.InitPostController(conn)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.GetFrontendURL()},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	r.RedirectTrailingSlash = true // 末尾スラッシュを自動リダイレクト
	authRouter := r.Group("/auth")
	postRouter := r.Group("/posts")
	postRouterWithAuth := r.Group("/posts", middlewares.AuthMiddleware(authUseCase))

	authRouter.POST("/signup", authController.SignUp)
	authRouter.POST("/login", authController.Login)

	postRouterWithAuth.POST("", postController.CreatePost)
	postRouterWithAuth.DELETE("/:did", postController.DeletePost)
	postRouterWithAuth.PUT("/:did", postController.UpdatePost)
	postRouterWithAuth.GET("/my", postController.GetPostsByUserID)
	postRouter.GET("", postController.GetPosts)
	postRouter.GET("/:did", postController.GetPostByDid)

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
