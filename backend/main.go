package main

import (
	"my-blog-backend/infra/db"

	_ "github.com/lib/pq" // Postgresドライバー
)

// func setUpRouter(db *sql.DB) *gin.Engine {
// 	// r := gin.Default()
// 	// userRouter := r.Group("/users")
// 	// authRouter := r.Group("/auth")
// }

func main() {
	db.Initialize()
	conn := db.SetupDB()
	defer conn.Close()

}
