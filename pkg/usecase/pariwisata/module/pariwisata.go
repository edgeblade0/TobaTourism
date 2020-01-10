package module

import (
	"log"
	"strconv"
	"fmt"

	"github.com/TobaTourism/pkg/models"
)

func (u *pariwisata) GetAllPariwisata() (models.PariwisataResponse, error) {
	allPariwisata, attachmentID, err := u.pariwisataRepo.GetAllPariwisata()
	if err != nil {
		log.Println("[Usecase][Pariwisata][GetAllPariwisata] Error: ", err)

		return allPariwisata, err
	}

	for i, id := range attachmentID {
		allPariwisata.Data[i].Attachment, err = u.attachmentRepo.GetAttachment(id)
		if err != nil {
			log.Println("[Restoran][Usecase][GetAttachment] Error : ", err)
			return allPariwisata, err
		}
	}
	return allPariwisata, nil
}

func (u *pariwisata) GetPariwisataByID(pariwisataID int64) (models.PariwisataResponse, error) {
	pariwisata, err := u.pariwisataRepo.GetPariwisataByID(pariwisataID)
	if err != nil {
		log.Println("[Usecase][Pariwisata][GetPariwisataByID] Error: ", err)

		return pariwisata, err
	}

	return pariwisata, nil
}

func (u *pariwisata) CreatePariwisata(nama string, lokasi string, description string, contact string, attachmentID int64) (models.PariwisataResponse, error) {
	pariwisata, err := u.pariwisataRepo.CreatePariwisata(nama, lokasi, description, contact, attachmentID)
	if err != nil {
		log.Println("[Usecase][Pariwisata][CreatePariwisata] Error: ", err)

		return pariwisata, err
	}

	return pariwisata, nil
}

func (u *pariwisata) UpdatePariwisata(pariwisataID int64, nama, lokasi, description, contact string) (models.PariwisataResponse, error) {
	pariwisata, err := u.pariwisataRepo.UpdatePariwisata(pariwisataID, nama, lokasi, description, contact)
	if err != nil {
		log.Println("[Usecase][Pariwisata][UpdatePariwisata] Error: ", err)

		return pariwisata, err
	}

	return pariwisata, nil
}

func (u *pariwisata) DeletePariwisata(pariwisataID int64) (models.PariwisataResponse, error) {
	pariwisata, err := u.pariwisataRepo.DeletePariwisata(pariwisataID)
	if err != nil {
		log.Println("[Usecase][Pariwisata][DeletePariwisata] Error: ", err)

		return pariwisata, err
	}

	return pariwisata, nil
}


func (k *pariwisata) UpdateImagePariwisata(pariwisataID string, attachmentID int64) error{
	var pariwisata models.Pariwisata
	
	pariwisataIDInt, err := strconv.ParseInt(pariwisataID, 10, 64)
	if err != nil {
		log.Println("[Usecase][Pariwisata][Parse pariwisatID on update Pariwisata] Error: ", err)
		return  err
	}

	pariwisata.ID = pariwisataIDInt
	

	if attachmentID == 0 {
		err = fmt.Errorf("[Usecase][Pariwisata][AttachmentID on update] no AttachmentID ")

		log.Println(err)
		return  err 
	}
	pariwisata.AttachmentID = attachmentID

	err = k.pariwisataRepo.UpdateImagePariwisata(pariwisata)
	if err != nil {
		log.Println("[Usecase][Pariwisata][UpdatePariwisata]", err)
		
		return   err
	}
	return  err
}


