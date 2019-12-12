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

	for i := range allExperience.Data {
		attachments, err := u.attachmentRepo.GetAttachment(allExperience.Data[i].AttachmentID)
		if err != nil {
			log.Println("[Usecase][Experience][GetAllExperience][GetAttachment] Error: ", err)

			return allExperience, err
		}

		allExperience.Data[i].Attachments = attachments
	}

	return allExperience, nil
}

func (u *experience) GetExperienceByID(experienceID int64) (models.ExperienceResponse, error) {
	experience, err := u.experienceRepo.GetExperienceByID(experienceID)
	if err != nil {
		log.Println("[Usecase][Experience][GetExperienceByID] Error: ", err)

		return experience, err
	}

	if experience.Data == nil || len(experience.Data) == 0 {
		log.Println("[Usecase][Experience][GetExperienceByID] Error: Data nil or not found")

		experience.Message = models.MessageNotFound
		experience.Status = models.StatusNotFound

		return experience, err
	}

	attachments, err := u.attachmentRepo.GetAttachment(experience.Data[0].AttachmentID)
	if err != nil {
		log.Println("[Usecase][Experience][GetExperienceByID][GetAttachment] Error: ", err)

		return experience, err
	}

	experience.Data[0].Attachments = attachments

	return experience, nil
}

func (u *experience) CreateExperience(description, location string, attachmentID int64) (models.ExperienceResponse, error) {
	experience, err := u.experienceRepo.CreateExperience(description, location, attachmentID)
	if err != nil {
		log.Println("[Usecase][Experience][CreateExperience] Error: ", err)

		return experience, err
	}

	return experience, nil
}

func (u *experience) UpdateExperience(experienceID int64, description, location string, attachmentID int64) (models.ExperienceResponse, error) {
	experience, err := u.experienceRepo.UpdateExperience(experienceID, description, location, attachmentID)
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
