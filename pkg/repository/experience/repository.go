package experience

import (
	"github.com/if416021/TobaTourism/pkg/models"
)

type Repository interface {
	GetAllExperience() (models.ExperienceResponse, error)
	GetExperienceByID(experienceID int64) (models.ExperienceResponse, error)
	CreateExperience(description, lokasi string) (models.ExperienceResponse, error)
	UpdateExperience(experienceID int64, description, lokasi string) (models.ExperienceResponse, error)
	DeleteExperience(experienceID int64) (models.ExperienceResponse, error)
}
