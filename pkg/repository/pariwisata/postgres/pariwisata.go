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

//for the insert
// func (r *pariwisata) CreatePariwisata() ([]models.Pariwisata, error) {
// 	allPariwisata := []models.Pariwisata{}

// 	rows, err := r.DB.Query(queryCreatePariwisata)
// 	// if err != nil {
// 	// 	log.Println("Repository error : ", err)
// 	// 	return allPariwisata, err
// 	// }

// 	// for rows.Next() {
// 	// 	pariwisata := models.Pariwisata{}
// 	// 	err := rows.Scan(&pariwisata.ID, &pariwisata.Name, &pariwisata.Lokasi)
// 	// 	if err != nil {
// 	// 		log.Println(err)
// 	// 	}
// 	// 	allPariwisata = append(allPariwisata, pariwisata)
// 	// }

// 	// return allPariwisata, nil
// }

func (r *pariwisata) CreatePariwisata(pariwisataName string, pariwisataLokasi string) error {

	statement, err := r.DB.Prepare(queryCreatePariwisata)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(pariwisataName, pariwisataLokasi)

	if err != nil {
		return err
	}
	return nil
}
