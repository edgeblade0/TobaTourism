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
