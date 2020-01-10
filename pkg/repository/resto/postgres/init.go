package postgres

import (
	"database/sql"

	restoRepo "github.com/if416021/TobaTourism/pkg/repository/resto"
)

type resto struct {
	DB *sql.DB
}

func InitRestoRepo(db *sql.DB) restoRepo.Repository {
	return &resto{
		DB: db,
	}
}
