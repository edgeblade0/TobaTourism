package postgres

import (
	"database/sql"

	attachmentRepo "github.com/if416021/TobaTourism/pkg/repository/attachment"
)

type attachment struct {
	DB *sql.DB
}

func InitAttachmentRepo(db *sql.DB) attachmentRepo.Repository {
	return &attachment{
		DB: db,
	}
}
