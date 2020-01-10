package module

import (
	"fmt"
	"log"
	"strconv"

	"github.com/TobaTourism/pkg/models"
)

func (u *resto) GetAllRestoWithKuliner() ([]models.Restoran, error) {
	var Restouran []models.Restoran
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

	for _, resto := range allResto {
		if len(resto.ListCulinary) > 0 {
			Restouran = append(Restouran, resto)
		}
	}

	return Restouran, err
}

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

	}

	return allResto, err
}

func (u *resto) GetDetailResto(restoID int64) (models.Restoran, error) {
	resto, err := u.restoRepo.GetDetailResto(restoID)
	if err != nil {
		log.Println("[Restoran][Usecase][GetDetailRestoran] Error : ", err)
		return resto, err
	}
	resto.Attachment, err = u.attachmentRepo.GetAttachment(resto.AttachmentID)
	if err != nil {
		log.Println("[Restoran][Usecase][GetAttachment] Error : ", err)
		return resto, err
	}

	resto.ListCulinary, err = u.kulinerUsecase.GetAllCulinary(resto.RestoID)
	if err != nil {
		log.Println("[Restoran][Usecase][GetAllKuliner] Error : ", err)
		return resto, err
	}

	// if len(resto.ListCulinary) == 0 {
	// 	resto.ListCulinary = new []models.Culinary
	// 	resto.ListCulinary[0].CulinaryID = 0
	// }

	// log.Println(resto.ListCulinary)

	return resto, err
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

func (u *resto) UpdateImageRestoran(restoID string, attachmentID int64) error {
	var resto models.Restoran

	restoIDInt, err := strconv.ParseInt(restoID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Restoran][Parse RestoID on Update] Error : ", err)
		return err
	}
	resto.RestoID = restoIDInt

	if attachmentID == 0 {
		err = fmt.Errorf("[Usecase][Restoran][AttachmentID] no AttachmentID")
		log.Println(err)
		return err
	}
	resto.AttachmentID = attachmentID

	err = u.restoRepo.UpdateImageResto(resto)
	if err != nil {
		log.Println("[Restoran][Usecase][CreateResto] Error : ", err)
		return err
	}

	return err
}

func (u *resto) UpdateRestoran(restoID string, restoName string, restoContact string, restoLocation string) error {
	var resto models.Restoran

	restoIDInt, err := strconv.ParseInt(restoID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Restoran][Parse RestoID on Update] Error : ", err)
		return err
	}
	resto.RestoID = restoIDInt

	if restoName == "" {
		err = fmt.Errorf("[Usecase][Restoran][restoName Update] no name")
		log.Println(err)
		return err
	}
	resto.Name = restoName

	if restoContact == "" {
		err = fmt.Errorf("[Usecase][Restoran][restoContact] no contact")
		log.Println(err)
		return err
	}
	resto.Contact = restoContact

	if restoLocation == "" {
		err = fmt.Errorf("[Usecase][Restoran][restoLocation] no location")
		log.Println(err)
		return err
	}
	resto.Location = restoLocation

	err = u.restoRepo.UpdateResto(resto)
	if err != nil {
		log.Println("[Restoran][Usecase][updateResto] Error : ", err)
		return err
	}

	return err
}

func (u *resto) DeleteResto(restoID string) error {
	restoIDInt, err := strconv.ParseInt(restoID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Restoran][Parse RestoID on Delete] Error : ", err)
		return err
	}

	err = u.restoRepo.DeleteResto(restoIDInt)
	if err != nil {
		log.Println("[Restoran][Usecase][DeleteResto] Error : ", err)
		return err
	}

	err = u.kulinerRepo.DeleteAllKuliner(restoIDInt)
	if err != nil {
		log.Println("[Restoran][Usecase][DeleteKuliner] Error : ", err)
		return err
	}

	return err
}
