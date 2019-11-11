package module

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (u *experience) GetAllExperience() (models.Response, error) {
	allExperience, err := u.experienceRepo.GetAllExperience()
	if err != nil {
		log.Println("[Usecase][Experience][GetAllExperience] Error: ", err)

		return allExperience, err
	}

	return allExperience, nil
}
