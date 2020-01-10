package postgres

import (
	"database/sql"

	pariwisataRepo "github.com/if416021/TobaTourism/pkg/repository/pariwisata"
)

type pariwisata struct {
	DB *sql.DB
}

func InitPariwisataRepo(db *sql.DB) pariwisataRepo.Repository {
	return &pariwisata{
		DB: db,
	}
}
