package module

import (
	"fmt"
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (u *resto) GetAllResto() ([]models.Restoran, error) {
	allResto, attachmentID, err := u.restoRepo.GetAllResto()
	if err != nil {
		log.Println("[Restoran][Usecase][GetAllRestoran] Error : ", err)
		return allResto, err
	}

	for i, id := range attachmentID {
		allResto[i].Attachment, err = u.attachmentRepo.GetAttachment(id)
		if err != nil {
			log.Println("[Restoran][Usecase][GetAttachment] Error : ", err)
			return allResto, err
		}

		allResto[i].ListCulinary, err = u.kulinerUsecase.GetAllCulinary(allResto[i].RestoID)
		if err != nil {
			log.Println("[Restoran][Usecase][GetAllKuliner] Error : ", err)
			return allResto, err
		}
	}

	return allResto, err
}

func (u *resto) CreateResto(name string, location string, contact string, attachmentID int64) error {
	var resto models.Restoran

	if name == "" {
		err := fmt.Errorf("[Restoran][Usecase][name] empty field name")
		log.Println(err)
		return err
	}
	resto.Name = name

	if location == "" {
		err := fmt.Errorf("[Restoran][Usecase][location] empty field location")
		log.Println(err)
		return err
	}
	resto.Location = location

	if contact == "" {
		err := fmt.Errorf("[Restoran][Usecase][contact] empty field contact")
		log.Println(err)
		return err
	}
	resto.Contact = contact

	if attachmentID == 0 {
		err := fmt.Errorf("[Restoran][Usecase][attachmentID] empty field attachmentID")
		log.Println(err)
		return err
	}
	resto.AttachmentID = attachmentID

	err := u.restoRepo.CreateResto(resto)
	if err != nil {
		log.Println("[Restoran][Usecase][CreateResto] Error : ", err)
		return err
	}

	return err
}
