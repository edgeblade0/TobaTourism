package postgres

import (
	"log"

	"github.com/TobaTourism/pkg/models"
)

const (
	MessageSucces = "Berhasil"
	Messagefailed = "Gagal"

	StatusOk     = "Ok"
	StatusFailed = "Gagal"
)

func (r *pariwisata) GetAllPariwisata() (models.Respone, error) {
	allPariwisata := []models.Pariwisata{}
	var resp models.Respone

	rows, err := r.DB.Query(QueryGetAllPariwisata)
	if err != nil {
		log.Println("[Repository][Experience][GetAllPariwisata] Query Error: ", err)

		resp.Data = allPariwisata
		resp.Message = Messagefailed
		resp.Status = StatusFailed

		return resp, err
	}

	for rows.Next() {
		pariwisata := models.Pariwisata{}
		err := rows.Scan(&pariwisata.ID, &pariwisata.Name, &pariwisata.Lokasi)
		if err != nil {
			log.Println("[Repository][Pariwisata][GetAllPariwisata] Scan Error : ", err)

			resp.Data = allPariwisata
			resp.Message = Messagefailed
			resp.Status = StatusFailed

			return resp, err
		}

		allPariwisata = append(allPariwisata, pariwisata)
	}

	resp.Data = allPariwisata
	resp.Message = MessageSucces
	resp.Status = StatusOk

	return resp, nil
}

func (r *pariwisata) CreatePariwisata(pariwisataName string, pariwisataLokasi string) error {

	statement, err := r.DB.Prepare(queryCreatePariwisata)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(pariwisataName, pariwisataLokasi)

	if err != nil {
		return err
	}
	return nil
}
