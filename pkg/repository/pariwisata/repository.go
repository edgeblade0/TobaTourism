package pariwisata

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	GetAllPariwisata() ([]models.Pariwisata, error)
	CreatePariwisata(pariwisataName, pariwisataLokasi string) error
}
