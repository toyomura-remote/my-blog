package infra

import "database/sql"

type userRepository struct {
	db *sql.DB
}
