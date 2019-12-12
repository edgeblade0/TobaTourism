package experience

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	GetAllExperience() (models.ExperienceResponse, error)
	GetExperienceByID(experienceID int64) (models.ExperienceResponse, error)
	CreateExperience(description, location string, attachmentID int64) (models.ExperienceResponse, error)
	UpdateExperience(experienceID int64, description, location string, attachmentID int64) (models.ExperienceResponse, error)
	DeleteExperience(experienceID int64) (models.ExperienceResponse, error)
}
