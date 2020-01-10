
package postgres

import (
	"log"

	"github.com/if416021/TobaTourism/pkg/models"
)

const (
	MessageSuccess = "Success"
	MessageFailed  = "Failed"

	StatusOK     = "200"
	StatusFailed = "500"
)

func (r *transportasi) GetAllTransportasi() (models.TransportasiResponse, error) {
	allTransportasi := []models.Transportasi{}
	var resp models.TransportasiResponse

	statement, err := r.DB.Prepare(QueryGetAllTransportasi)
	if err != nil {
		log.Println("[Repository][Transportasi][GetAllTransportasi] Prepare error: ", err)

		resp.Data = allTransportasi
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		log.Println("[Repository][Transportasi][GetAllTransportasi] Query error: ", err)

		resp.Data = allTransportasi
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	for rows.Next() {
		transportasi := models.Transportasi{}
		err := rows.Scan(&transportasi.ID, &transportasi.Nama, &transportasi.Rute, &transportasi.Description, &transportasi.Contact, &transportasi.Harga)
		if err != nil {
			log.Println("[Repository][Transportasi][GetAllTransportasi] Scan error: ", err)

			resp.Data = allTransportasi
			resp.Message = MessageFailed
			resp.Status = StatusFailed

			return resp, err
		}

		allTransportasi = append(allTransportasi, transportasi)
	}

	resp.Data = allTransportasi
	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *transportasi) GetTransportasiByID(transportasiID int64) (models.TransportasiResponse, error) {
	transportasis := []models.Transportasi{}
	var resp models.TransportasiResponse

	statement, err := r.DB.Prepare(QueryGetTransportasiByID)
	if err != nil {
		log.Println("[Repository][Transportasi][GetTransportasiByID] Prepare error: ", err)

		resp.Data = transportasis
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	rows, err := statement.Query(transportasiID)
	if err != nil {
		log.Println("[Repository][Transportasi][GetTransportasiByID] Query error: ", err)

		resp.Data = transportasis
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	for rows.Next() {
		transportasi := models.Transportasi{}
		err = rows.Scan(&transportasi.ID, &transportasi.Nama, &transportasi.Rute, &transportasi.Description, &transportasi.Contact, &transportasi.Harga)
		if err != nil {
			log.Println("[Repository][Transportasi][GetTransportasiByID] Scan error: ", err)

			resp.Data = transportasis
			resp.Message = MessageFailed
			resp.Status = StatusFailed

			return resp, err
		}

		transportasis = append(transportasis, transportasi)
	}

	resp.Data = transportasis
	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *transportasi) CreateTransportasi(nama, rute, description, contact string, harga int64) (models.TransportasiResponse, error) {
	var resp models.TransportasiResponse

	statement, err := r.DB.Prepare(QueryCreateTransportasi)
	if err != nil {
		log.Println("[Repository][Transportasi][CreateTransportasi] Prepare error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	_, err = statement.Exec(nama, rute, description, contact, harga)
	if err != nil {
		log.Println("[Repository][Transportasi][CreateTransportasi] Exec error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *transportasi) UpdateTransportasi(transportasiID int64, nama, rute, description, contact string, harga int64) (models.TransportasiResponse, error) {
	var resp models.TransportasiResponse

	statement, err := r.DB.Prepare(QueryUpdateTransportasi)
	if err != nil {
		log.Println("[Repository][Transportasi][UpdateTransportasi] Prepare error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	_, err = statement.Exec(transportasiID, nama, rute, description, contact, harga)
	if err != nil {
		log.Println("[Repository][Transportasi][UpdateTransportasi] Exec error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *transportasi) DeleteTransportasi(transportasiID int64) (models.TransportasiResponse, error) {
	var resp models.TransportasiResponse

	statement, err := r.DB.Prepare(QueryDeleteTransportasi)
	if err != nil {
		log.Println("[Repository][Transportasi][DeleteTransportasi] Prepare error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	_, err = statement.Exec(transportasiID)
	if err != nil {
		log.Println("[Repository][Transportasi][DeleteTransportasi] Exec error:", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed
	}

	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}