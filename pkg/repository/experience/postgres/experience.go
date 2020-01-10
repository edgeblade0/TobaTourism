package postgres

import (
	"log"

	"github.com/if416021/TobaTourism/pkg/models"
)

const (
	MessageSuccess = "Success"
	MessageFailed  = "Failed"

	StatusOK     = "200"
	StatusFailed = "500"
)

func (r *experience) GetAllExperience() (models.ExperienceResponse, error) {
	allExperience := []models.Experience{}
	var resp models.ExperienceResponse

	statement, err := r.DB.Prepare(QueryGetAllExperience)
	if err != nil {
		log.Println("[Repository][Experience][GetAllExperience] Prepare error: ", err)

		resp.Data = allExperience
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		log.Println("[Repository][Experience][GetAllExperience] Query error: ", err)

		resp.Data = allExperience
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	for rows.Next() {
		experience := models.Experience{}
		err := rows.Scan(&experience.ID, &experience.Description, &experience.Lokasi)
		if err != nil {
			log.Println("[Repository][Experience][GetAllExperience] Scan error: ", err)

			resp.Data = allExperience
			resp.Message = MessageFailed
			resp.Status = StatusFailed

			return resp, err
		}

		allExperience = append(allExperience, experience)
	}

	resp.Data = allExperience
	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *experience) GetExperienceByID(experienceID int64) (models.ExperienceResponse, error) {
	experiences := []models.Experience{}
	var resp models.ExperienceResponse

	statement, err := r.DB.Prepare(QueryGetExperienceByID)
	if err != nil {
		log.Println("[Repository][Experience][GetExperienceByID] Prepare error: ", err)

		resp.Data = experiences
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	rows, err := statement.Query(experienceID)
	if err != nil {
		log.Println("[Repository][Experience][GetExperienceByID] Query error: ", err)

		resp.Data = experiences
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	for rows.Next() {
		experience := models.Experience{}
		err = rows.Scan(&experience.ID, &experience.Description, &experience.Lokasi)
		if err != nil {
			log.Println("[Repository][Experience][GetExperienceByID] Scan error: ", err)

			resp.Data = experiences
			resp.Message = MessageFailed
			resp.Status = StatusFailed

			return resp, err
		}

		experiences = append(experiences, experience)
	}

	resp.Data = experiences
	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *experience) CreateExperience(description, lokasi string) (models.ExperienceResponse, error) {
	var resp models.ExperienceResponse

	statement, err := r.DB.Prepare(QueryCreateExperience)
	if err != nil {
		log.Println("[Repository][Experience][CreateExperience] Prepare error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	_, err = statement.Exec(description, lokasi)
	if err != nil {
		log.Println("[Repository][Experience][CreateExperience] Exec error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *experience) UpdateExperience(experienceID int64, description, lokasi string) (models.ExperienceResponse, error) {
	var resp models.ExperienceResponse

	statement, err := r.DB.Prepare(QueryUpdateExperience)
	if err != nil {
		log.Println("[Repository][Experience][UpdateExperience] Prepare error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	_, err = statement.Exec(experienceID, description, lokasi)
	if err != nil {
		log.Println("[Repository][Experience][UpdateExperience] Exec error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *experience) DeleteExperience(experienceID int64) (models.ExperienceResponse, error) {
	var resp models.ExperienceResponse

	statement, err := r.DB.Prepare(QueryDeleteExperience)
	if err != nil {
		log.Println("[Repository][Experience][DeleteExperience] Prepare error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	_, err = statement.Exec(experienceID)
	if err != nil {
		log.Println("[Repository][Experience][DeleteExperience] Exec error:", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed
	}

	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}
