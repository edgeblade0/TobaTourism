package module

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (u *profile) GetProfile(userID int64) (models.Profile, error) {
	var profile models.Profile

	profile, err := u.profileRepo.GetProfile(userID)

	if err != nil {
		log.Println("[Kuliner][Usecase][GetDetailKuliner] Error : ", err)
		return profile, err
	}

	return profile, nil
}

func (u *profile) UpdateProfile(profile models.Profile) error {

	return nil
}
