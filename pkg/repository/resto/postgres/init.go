package postgres

import (
	"database/sql"

	restoRepo "github.com/TobaTourism/pkg/repository/resto"
)

type resto struct {
	DB *sql.DB
}

func InitRestoRepo(db *sql.DB) restoRepo.Repository {
	return &resto{
		DB: db,
	}
}
