package postgres

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

const (
	MessageSuccess = "Success"
	MessageFailed  = "Failed"

	StatusOK     = "200"
	StatusFailed = "500"
)

func (r *pariwisata) GetAllPariwisata() (models.PariwisataResponse, error) {
	allPariwisata := []models.Pariwisata{}
	var resp models.PariwisataResponse

	statement, err := r.DB.Prepare(QueryGetAllPariwisata)
	if err != nil {
		log.Println("[Repository][Pariwisata][GetAllPariwisata] Prepare error: ", err)

		resp.Data = allPariwisata
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		log.Println("[Repository][Pariwisata][GetAllPariwisata] Query error: ", err)

		resp.Data = allPariwisata
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	for rows.Next() {
		pariwisata := models.Pariwisata{}
		err := rows.Scan(&pariwisata.ID, &pariwisata.Nama, &pariwisata.Lokasi, &pariwisata.Description, &pariwisata.Contact)
		if err != nil {
			log.Println("[Repository][Pariwisata][GetAllPariwisata] Scan error: ", err)

			resp.Data = allPariwisata
			resp.Message = MessageFailed
			resp.Status = StatusFailed

			return resp, err
		}

		allPariwisata = append(allPariwisata, pariwisata)
	}

	resp.Data = allPariwisata
	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *pariwisata) GetPariwisataByID(pariwisataID int64) (models.PariwisataResponse, error) {
	pariwisatas := []models.Pariwisata{}
	var resp models.PariwisataResponse

	statement, err := r.DB.Prepare(QueryGetPariwisataByID)
	if err != nil {
		log.Println("[Repository][Pariwisata][GetPariwisataByID] Prepare error: ", err)

		resp.Data = pariwisatas
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	rows, err := statement.Query(pariwisataID)
	if err != nil {
		log.Println("[Repository][Pariwisata][GetPariwisataByID] Query error: ", err)

		resp.Data = pariwisatas
		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	for rows.Next() {
		pariwisata := models.Pariwisata{}
		err = rows.Scan(&pariwisata.ID, &pariwisata.Nama, &pariwisata.Lokasi, &pariwisata.Description, &pariwisata.Contact)
		if err != nil {
			log.Println("[Repository][Pariwisata][GetPariwisataByID] Scan error: ", err)

			resp.Data = pariwisatas
			resp.Message = MessageFailed
			resp.Status = StatusFailed

			return resp, err
		}

		pariwisatas = append(pariwisatas, pariwisata)
	}

	resp.Data = pariwisatas
	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *pariwisata) CreatePariwisata(nama, lokasi, description, contact string) (models.PariwisataResponse, error) {
	var resp models.PariwisataResponse
	log.Println(nama)
	statement, err := r.DB.Prepare(QueryCreatePariwisata)
	if err != nil {
		log.Println("[Repository][Pariwisata][CreatePariwisata] Prepare error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	_, err = statement.Exec(nama, lokasi, description, contact)
	if err != nil {
		log.Println("[Repository][Pariwisata][CreatePariwisata] Exec error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *pariwisata) UpdatePariwisata(pariwisataID int64, nama, lokasi, description, contact string) (models.PariwisataResponse, error) {
	var resp models.PariwisataResponse

	statement, err := r.DB.Prepare(QueryUpdatePariwisata)
	if err != nil {
		log.Println("[Repository][Pariwisata][UpdatePariwisata] Prepare error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	_, err = statement.Exec(pariwisataID, nama, lokasi, description, contact)
	if err != nil {
		log.Println("[Repository][Pariwisata][UpdatePariwisata] Exec error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}

func (r *pariwisata) DeletePariwisata(pariwisataID int64) (models.PariwisataResponse, error) {
	var resp models.PariwisataResponse

	statement, err := r.DB.Prepare(QueryDeletePariwisata)
	if err != nil {
		log.Println("[Repository][Pariwisata][DeletePariwisata] Prepare error: ", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed

		return resp, err
	}

	defer statement.Close()

	_, err = statement.Exec(pariwisataID)
	if err != nil {
		log.Println("[Repository][Pariwisata][DeletePariwisata] Exec error:", err)

		resp.Message = MessageFailed
		resp.Status = StatusFailed
	}

	resp.Message = MessageSuccess
	resp.Status = StatusOK

	return resp, nil
}
