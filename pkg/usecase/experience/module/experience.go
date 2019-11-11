package module

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (u *experience) GetAllExperience() ([]models.Experience, error) {
	allExperience, err := u.experienceRepo.GetAllExperience()
	if err != nil {
		log.Println(err)
		return allExperience, err
	}

	return allExperience, err
}
