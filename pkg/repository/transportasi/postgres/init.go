package postgres

import (
	"database/sql"

	transportasiRepo "github.com/if416021/TobaTourism/pkg/repository/transportasi"
)

type transportasi struct {
	DB *sql.DB
}

// InitTransportasiRepo to initialize repository
func InitTransportasiRepo(db *sql.DB) transportasiRepo.Repository {
	return &transportasi{
		DB: db,
	}
}
