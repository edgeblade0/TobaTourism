package postgres

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (r *pariwisata) GetAllPariwisata() ([]models.Pariwisata, error) {
	allPariwisata := []models.Pariwisata{}

	rows, err := r.DB.Query(queryGetAllPariwisata)
	if err != nil {
		log.Println("Repository error : ", err)
		return allPariwisata, err
	}

	for rows.Next() {
		pariwisata := models.Pariwisata{}
		err := rows.Scan(&pariwisata.ID, &pariwisata.Name, &pariwisata.Lokasi)
		if err != nil {
			log.Println(err)
		}
		allPariwisata = append(allPariwisata, pariwisata)
	}

	return allPariwisata, nil
}
