package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq" // Postgresドライバー
)

func main() {
    // 接続情報（docker-composeで設定した名前を使う！）
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    // DBに接続
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 実際に繋がるか確認（ピンポン打診）
    err = db.Ping()
    if err != nil {
        log.Fatal("DB接続失敗...: ", err)
    }

    fmt.Println("DB接続成功！GoとPostgresが友達になりました！")
}