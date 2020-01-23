package postgres

import (
	"database/sql"

	profileRepo "github.com/TobaTourism/pkg/repository/profile"
)

type profile struct {
	DB *sql.DB
}

func InitProfileRepo(db *sql.DB) profileRepo.Repository {
	return &profile{
		DB: db,
	}
}
