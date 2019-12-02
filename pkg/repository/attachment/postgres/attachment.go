package postgres

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (a *attachment) InsertFile(Attachment models.Attachment) (int64, error) {
	var AttachmentID int64
	statement, err := a.DB.Prepare(QueryInsertAttachment)
	if err != nil {
		log.Println("[Repository][Attachment][Prepare Insert] Error : ", err)
		return AttachmentID, err
	}
	defer statement.Close()

	err = statement.QueryRow(Attachment.Type, Attachment.Content).Scan(&AttachmentID)
	if err != nil {
		log.Println("[Repository][Attachmen][Execute] Error : ", err)
		return AttachmentID, err
	}
	return AttachmentID, nil
}

func (a *attachment) GetAttachment(AttchmentID int64) (string, error) {
	var Content string
	statement, err := a.DB.Prepare(QuerySelectAttachment)
	if err != nil {
		log.Println("[Repository][Attachment][Prepare GetAttachment] Error : ", err)
		return Content, err
	}

	statement.QueryRow(AttchmentID).Scan(&Content)

	return Content, err

}
