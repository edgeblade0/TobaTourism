package pariwisata

import "github.com/TobaTourism/pkg/models"

type Usecase interface {
	GetAllPariwisata() (models.PariwisataResponse, error)
	GetPariwisataByID(pariwisataID int64) (models.PariwisataResponseByID, error)
	CreatePariwisata(nama string, lokasi string, description string, contact string, attachmentID int64) (models.PariwisataResponse, error)
	UpdatePariwisata(pariwisataID int64, nama, lokasi, description, contact string) (models.PariwisataResponse, error)
	DeletePariwisata(pariwisataID int64) (models.PariwisataResponse, error)
	UpdateImagePariwisata(pariwisataID string, attachemntID int64) error
}
