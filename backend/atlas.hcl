env "local" {
    url = getenv("DB_URL_LOCAL")
    dev = getenv("DB_URL_LOCAL_DEV")

    migration {
        dir             = "file://db/migrations"
        format          = atlas
        revisions_schema = "public"
    }

    src = "file://db/original"
}

env "prod" {
    url = getenv("DB_URL_PROD")
    dev = getenv("DB_URL_PROD_DEV")

    migration {
        dir = "file://db/migrations"
    }
}