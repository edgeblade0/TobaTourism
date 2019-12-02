package module

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (u *pariwisata) GetAllPariwisata() (models.PariwisataResponse, error) {
	allPariwisata, err := u.pariwisataRepo.GetAllPariwisata()
	if err != nil {
		log.Println("[Usecase][Pariwisata][GetAllPariwisata] Error: ", err)

		return allPariwisata, err
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

func (u *pariwisata) CreatePariwisata(nama, lokasi, description, contact string) (models.PariwisataResponse, error) {
	pariwisata, err := u.pariwisataRepo.CreatePariwisata(nama, lokasi, description, contact)
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
