package experience

import (
	"github.com/TobaTourism/pkg/models"
)

// Usecase interface
type Usecase interface {
	GetAllExperience() (models.Response, error)
}
