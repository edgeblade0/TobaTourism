package postgres

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (r *resto) GetAllResto() ([]models.Restoran, []int64, error) {
	allResto := []models.Restoran{}
	var AttachmentID []int64

	statement, err := r.DB.Prepare(QuerySelectRestoran)
	if err != nil {
		log.Println("[Repository][Restoran][Prepare] Error : ", err)
		return allResto, AttachmentID, err
	}

	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		log.Println("[Repository][Restoran][Query] Error : ", err)
		return allResto, AttachmentID, err
	}

	for rows.Next() {
		var restoran models.Restoran
		var attachID int64
		err := rows.Scan(&restoran.RestoID, &restoran.Contact, &restoran.Location, &restoran.Name, &attachID)
		if err != nil {
			log.Println("[Repository][Restoran][Row Next] Error : ", err)
			return allResto, AttachmentID, err
		}
		AttachmentID = append(AttachmentID, attachID)
		allResto = append(allResto, restoran)
	}

	return allResto, AttachmentID, err
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

func (r *resto) UpdateImageResto(resto models.Restoran) error {
	statement, err := r.DB.Prepare(QueryUpdateImageRestoran)
	if err != nil {
		log.Println("[Repository][Restoran][Prepare Update Image] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(resto.AttachmentID, resto.RestoID)
	if err != nil {
		log.Println("[Repository][Restoran][Execute Update Image] Error : ", err)
		return err
	}
	return err
}

func (r *resto) UpdateResto(resto models.Restoran) error {
	log.Println(resto)
	statement, err := r.DB.Prepare(QueryUpdateRestoran)
	if err != nil {
		log.Println("[Repository][Restoran][Prepare Update] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(resto.Contact, resto.Location, resto.Name, resto.RestoID)
	if err != nil {
		log.Println("[Repository][Restoran][Execute Update] Error : ", err)
		return err
	}
	return err
}

func (r *resto) DeleteResto(restoID int64) error {
	statement, err := r.DB.Prepare(QueryDeleteRestoran)
	if err != nil {
		log.Println("[Repository][Restoran][Prepare delete] Error : ", err)
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(restoID)
	if err != nil {
		log.Println("[Repository][Restoran][Execute delete] Error : ", err)
		return err
	}
	return err
}
