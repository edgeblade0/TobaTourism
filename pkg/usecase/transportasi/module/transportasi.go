package module

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

func (u *transportasi) GetAllTransportasi() (models.TransportasiResponse, error) {
	allTransportasi, err := u.transportasiRepo.GetAllTransportasi()
	if err != nil {
		log.Println("[Usecase][Transportasi][GetAllTransportasi] Error: ", err)

		return allTransportasi, err
	}

	return allTransportasi, nil
}

func (u *transportasi) GetTransportasiByID(transportasiID int64) (models.TransportasiResponse, error) {
	transportasi, err := u.transportasiRepo.GetTransportasiByID(transportasiID)
	if err != nil {
		log.Println("[Usecase][Transportasi][GetTransportasiByID] Error: ", err)

		return transportasi, err
	}

	return transportasi, nil
}

func (u *transportasi) CreateTransportasi(nama, rute, description, contact string, harga int64) (models.TransportasiResponse, error) {
	transportasi, err := u.transportasiRepo.CreateTransportasi(nama, rute, description, contact, harga)
	if err != nil {
		log.Println("[Usecase][Transportasi][CreateTransportasi] Error: ", err)

		return transportasi, err
	}

	return transportasi, nil
}

func (u *transportasi) UpdateTransportasi(transportasiID int64, nama, rute, description, contact string, harga int64) (models.TransportasiResponse, error) {
	transportasi, err := u.transportasiRepo.UpdateTransportasi(transportasiID, nama, rute, description, contact, harga)
	if err != nil {
		log.Println("[Usecase][Transportasi][UpdateTransportasi] Error: ", err)

		return transportasi, err
	}

	return transportasi, nil
}

func (u *transportasi) DeleteTransportasi(transportasiID int64) (models.TransportasiResponse, error) {
	transportasi, err := u.transportasiRepo.DeleteTransportasi(transportasiID)
	if err != nil {
		log.Println("[Usecase][Transportasi][DeleteTransportasi] Error: ", err)

		return transportasi, err
	}

	return transportasi, nil
}
