package postgres

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

const (
	MessageSuccess = "Berhasil"
	MessageFailed  = "Gagal"

	StatusOK     = "OK"
	StatusFailed = "Failed"
)

func (r *experience) GetAllExperience() (models.Response, error) {
	allExperience := []models.Experience{}
	var resp models.Response

	rows, err := r.DB.Query(QueryGetAllExperience)
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
