package postgres

import (
	"database/sql"

	kulinerRepo "github.com/if416021/TobaTourism/pkg/repository/kuliner"
)

type kuliner struct {
	DB *sql.DB
}

func InitKulinerRepo(db *sql.DB) kulinerRepo.Repository {
	return &kuliner{
		DB: db,
	}
}
