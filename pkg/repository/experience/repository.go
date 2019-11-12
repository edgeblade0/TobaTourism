package experience

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	GetAllExperience() (models.Response, error)
}
