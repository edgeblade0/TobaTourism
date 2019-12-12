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

func (k *kuliner) GetDetailKuliner(RestoID int64, CulinaryID int64) (models.Culinary, int64, error) {
	var culinary models.Culinary
	var attachmentID int64

	statement, err := k.DB.Prepare(QuerySelectDetailCulinary)
	if err != nil {
		log.Println("[Repository][Kuliner][Prepare Detail] Error : ", err)
		return culinary, attachmentID, err
	}

	defer statement.Close()

	rows, err := statement.Query(RestoID, CulinaryID)
	if err != nil {
		log.Println("[Repository][Kuliner][Query Detail] Error : ", err)
		return culinary, attachmentID, err
	}

	for rows.Next() {
		err := rows.Scan(&culinary.Name, &culinary.Price, &culinary.Desc, &attachmentID)
		if err != nil {
			log.Println("[Repository][Kuliner][Row Next Detail] Error : ", err)
			return culinary, attachmentID, err
		}
	}

	return culinary, attachmentID, nil
}

func (r *kuliner) UpdateKuliner(kuliner models.Culinary) error {
	statement, err := r.DB.Prepare(QueryUpdateCulinary)
	if err != nil {
		log.Println("[Repository][Kuliner][Prepare Update] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(kuliner.Name, kuliner.Price, kuliner.Desc, kuliner.RestoID, kuliner.CulinaryID)
	if err != nil {
		log.Println("[Repository][Kuliner][Execute Update] Error : ", err)
		return err
	}
	return err
}

func (r *kuliner) UpdateImageKuliner(kuliner models.Culinary) error {
	statement, err := r.DB.Prepare(QueryUpdateImageCulinary)
	if err != nil {
		log.Println("[Repository][Kuliner][Prepare Update] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(kuliner.AttachmentID, kuliner.RestoID, kuliner.CulinaryID)
	if err != nil {
		log.Println("[Repository][Kuliner][Execute Update] Error : ", err)
		return err
	}
	return err
}

func (r *kuliner) DeleteKuliner(kuliner models.Culinary) error {
	statement, err := r.DB.Prepare(QueryDeleteCulinary)
	if err != nil {
		log.Println("[Repository][Kuliner][Prepare Delete] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(kuliner.RestoID, kuliner.CulinaryID)
	if err != nil {
		log.Println("[Repository][Kuliner][Execute Delete] Error : ", err)
		return err
	}
	return err
}

func (r *kuliner) DeleteAllKuliner(restoID int64) error {
	statement, err := r.DB.Prepare(QueryDeleteAllCulinary)
	if err != nil {
		log.Println("[Repository][Kuliner][Prepare DeleteAll] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(restoID)
	if err != nil {
		log.Println("[Repository][Kuliner][Execute DeleteAll] Error : ", err)
		return err
	}
	return err
}
