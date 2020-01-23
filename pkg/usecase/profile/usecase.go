package profile

import "github.com/TobaTourism/pkg/models"

type Usecase interface {
	GetProfile(userID int64) (models.Profile, error)
	UpdateProfile(profile models.Profile) error
}
