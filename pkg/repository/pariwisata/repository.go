package pariwisata

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	GetAllPariwisata() (models.PariwisataResponse, []int64, error)
	GetPariwisataByID(pariwisataID int64) (models.PariwisataResponseByID, error)
	CreatePariwisata(nama, lokasi, description, contact string, attachmentID int64) (models.PariwisataResponse, error)
	UpdatePariwisata(pariwisataID int64, nama, lokasi, description, contact string) (models.PariwisataResponse, error)
	DeletePariwisata(pariwisataID int64) (models.PariwisataResponse, error)
	UpdateImagePariwisata(pariwisata models.Pariwisata) error
}
