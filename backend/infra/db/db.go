package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func SetupDB() *sql.DB {
	// 接続情報
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// DB接続
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("DB Connect fatal: ", err)
	}

	return db
}
