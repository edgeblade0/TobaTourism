package postgres

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (r *resto) GetAllResto() ([]models.Restoran, error) {
	allResto := []models.Restoran{}

	// rows, err := r.DB.Query(queryGetAllResto)
	// if err != nil {
	// 	log.Println("Repository error : ", err)
	// 	return allResto, err
	// }

	// for rows.Next() {
	// 	pariwisata := models.Resto{}
	// 	err := rows.Scan(&pariwisata.ID, &pariwisata.Name, &pariwisata.Lokasi)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	allResto = append(allPariwisata, pariwisata)
	// }

	return allResto, nil
}

func (r *resto) CreateResto(resto models.Restoran) error {
	statement, err := r.DB.Prepare(QueryInsertRestoran)
	if err != nil {
		log.Println("[Repository][Restoran][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(resto.Contact, resto.Location, resto.Name, resto.AttachmentID)
	if err != nil {
		log.Println("[Repository][Restoran][Execute] Error : ", err)
		return err
	}
	return err
}
