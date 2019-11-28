package module

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (u *experience) GetAllExperience() (models.ExperienceResponse, error) {
	allExperience, err := u.experienceRepo.GetAllExperience()
	if err != nil {
		log.Println("[Usecase][Experience][GetAllExperience] Error: ", err)

		return allExperience, err
	}

	return allExperience, nil
}

func (u *experience) GetExperienceByID(experienceID int64) (models.ExperienceResponse, error) {
	experience, err := u.experienceRepo.GetExperienceByID(experienceID)
	if err != nil {
		log.Println("[Usecase][Experience][GetExperienceByID] Error: ", err)

		return experience, err
	}

	return experience, nil
}

func (u *experience) CreateExperience(description, lokasi string) (models.ExperienceResponse, error) {
	experience, err := u.experienceRepo.CreateExperience(description, lokasi)
	if err != nil {
		log.Println("[Usecase][Experience][CreateExperience] Error: ", err)

		return experience, err
	}

	return experience, nil
}

func (u *experience) UpdateExperience(experienceID int64, description, lokasi string) (models.ExperienceResponse, error) {
	experience, err := u.experienceRepo.UpdateExperience(experienceID, description, lokasi)
	if err != nil {
		log.Println("[Usecase][Experience][UpdateExperience] Error: ", err)

		return experience, err
	}

	return experience, nil
}

func (u *experience) DeleteExperience(experienceID int64) (models.ExperienceResponse, error) {
	experience, err := u.experienceRepo.DeleteExperience(experienceID)
	if err != nil {
		log.Println("[Usecase][Experience][DeleteExperience] Error: ", err)

		return experience, err
	}

	return experience, nil
}
