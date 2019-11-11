package postgres

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (r *experience) GetAllExperience() ([]models.Experience, error) {
	allExperience := []models.Experience{}

	rows, err := r.DB.Query(queryGetAllExperience)
	if err != nil {
		log.Println("Repository error: ", err)
		return allExperience, err
	}

	for rows.Next() {
		experience := models.Experience{}
		err := rows.Scan(&experience.ID, &experience.Description, &experience.Lokasi)
		if err != nil {
			log.Println(err)
		}

		allExperience = append(allExperience, experience)
	}

	return allExperience, nil
}
