env "local" {
    url = "postgres://user:password@db:5432/blog_db?sslmode=disable"
    dev = "postgres://user:password@db:5432/atlas_dev?sslmode=disable"

    migration {
        dir = "file://db/migrations"
        format = atlas
    }

    src = "file://db/original"
}