package attachment

import (
	"github.com/if416021/TobaTourism/pkg/models"
)

type Repository interface {
	InsertFile(attachment models.Attachment) (int64, error)
	GetAttachment(AttachmentID int64) ([]string, error)
}
