package postgres

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (r *kuliner) CreateKuliner(culinary models.Culinary) error {
	statement, err := r.DB.Prepare(QueryInsertKuliner)
	if err != nil {
		log.Println("[Repository][Kuliner][Prepare] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(culinary.RestoID, culinary.Name, culinary.Price, culinary.Desc, culinary.AttachmentID)
	if err != nil {
		log.Println("[Repository][Kuliner][Execute] Error : ", err)
		return err
	}
	return err
}

func (r *kuliner) GetAllKuliner(RestoID int64) ([]models.Culinary, []int64, error) {
	var allCulinary []models.Culinary
	var attachmnentID []int64
	statement, err := r.DB.Prepare(QuerySelectAllCulinary)
	if err != nil {
		log.Println("[Repository][Kuliner][Prepare] Error : ", err)
		return allCulinary, attachmnentID, err
	}

	defer statement.Close()

	rows, err := statement.Query(RestoID)
	if err != nil {
		log.Println("[Repository][Kuliner][Query] Error : ", err)
		return allCulinary, attachmnentID, err
	}

	for rows.Next() {
		var culinary models.Culinary
		var attachID int64
		err := rows.Scan(&culinary.CulinaryID, &culinary.Name, &culinary.Price, &culinary.Desc, &attachID)
		if err != nil {
			log.Println("[Repository][Kuliner][Row Next] Error : ", err)
			return allCulinary, attachmnentID, err
		}
		attachmnentID = append(attachmnentID, attachID)
		allCulinary = append(allCulinary, culinary)
	}
	return allCulinary, attachmnentID, err
}
