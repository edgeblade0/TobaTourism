package pariwisata

import "github.com/if416021/TobaTourism/pkg/models"

type Usecase interface {
	GetAllPariwisata() (models.PariwisataResponse, error)
	GetPariwisataByID(pariwisataID int64) (models.PariwisataResponse, error)
	CreatePariwisata(nama, lokasi, description, contact string) (models.PariwisataResponse, error)
	UpdatePariwisata(pariwisataID int64, nama, lokasi, description, contact string) (models.PariwisataResponse, error)
	DeletePariwisata(pariwisataID int64) (models.PariwisataResponse, error)
}
