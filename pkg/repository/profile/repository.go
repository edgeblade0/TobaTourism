package profile

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	GetProfile(userID int64) (models.Profile, error)
	UpdateProfile(profile models.Profile) error
}
