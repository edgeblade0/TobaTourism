package postgres

import (
	"database/sql"

	experienceRepo "github.com/if416021/TobaTourism/pkg/repository/experience"
)

type experience struct {
	DB *sql.DB
}

// InitExperienceRepo to initialize repository
func InitExperienceRepo(db *sql.DB) experienceRepo.Repository {
	return &experience{
		DB: db,
	}
}
