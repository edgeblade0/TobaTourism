package module

import (
	"fmt"
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (u *resto) GetAllResto() ([]models.Restoran, error) {
	allResto, err := u.restoRepo.GetAllResto()
	if err != nil {
		log.Println(err)
		return allResto, err
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

	resto.Location = location
	resto.Contact = contact
	resto.AttachmentID = attachmentID

	err := u.restoRepo.CreateResto(resto)
	if err != nil {
		log.Println("[Restoran][Usecase][CreateResto] Error : ", err)
		return err
	}

	return err
}
