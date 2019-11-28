package module

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/TobaTourism/pkg/models"
)

func (u *attachment) InsertAttachment(files []*multipart.FileHeader, FilePath string, Type int) (int64, error) {
	var AttachmentID int64
	var pathFile []string

	if len(files) == 0 {
		err := fmt.Errorf("[Usecase][Attachment][LenghtFile] No File")
		return AttachmentID, err
	}
	for _, file := range files {
		source, err := file.Open()
		if err != nil {
			log.Println("[attachment][usecase][OpenFile] Error : ", err)
			return AttachmentID, err
		}
		defer source.Close()

		path := FilePath + file.Filename
		destination, err := os.Create(path)
		if err != nil {
			log.Println("[attachment][usecase][CreateFile] Error : ", err)
			return AttachmentID, err
		}

		path = models.HomeAttachment + path
		pathFile = append(pathFile, path)

		defer destination.Close()

		if _, err = io.Copy(destination, source); err != nil {
			log.Println("[attachment][usecase][CopyFile] Error : ", err)
			return AttachmentID, err
		}
	}
	var attachment models.Attachment
	attachment.Type = Type
	bytePathFile, err := json.Marshal(pathFile)
	if err != nil {
		log.Println("Error Marshal [usecase][attachment][Marshal] Error: ", err)
		return AttachmentID, err
	}
	attachment.Content = string(bytePathFile)
	AttachmentID, err = u.attachmentRepo.InsertFile(attachment)
	if err != nil {
		log.Println("[attachment][usecase][InsertFile] Error : ", err)
		return AttachmentID, err
	}

	return AttachmentID, err
}

// func (u *attachment) GetAttachment(AttachmentID string) error {
// 	AttachmentIDInt, err := strconv.ParseInt(AttachmentID, 10, 64)
// 	if err != nil {
// 		log.Println("[Usecase][Attachment][Parse AttachmentID] Error : ", err)
// 		return nil
// 	}
// 	contentAttachment, err := u.attachmentRepo.GetAttachment(AttachmentIDInt)
// 	if err != nil {
// 		log.Println("[Usecase][Attachment][GetAttachment] Error : ", err)
// 		return nil
// 	}

// 	log.Println("Content : ", contentAttachment)

// 	return err
// }
