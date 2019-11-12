package experience

import (
	"github.com/TobaTourism/pkg/models"
)

// Usecase interface
type Usecase interface {
	GetAllExperience() (models.Response, error)
	GetExperienceByID(experienceID int64) (models.Response, error)
	CreateExperience(description, lokasi string) (models.Response, error)
	UpdateExperience(experienceID int64, description, lokasi string) (models.Response, error)
	DeleteExperience(experienceID int64) (models.Response, error)
}
