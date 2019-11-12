package pariwisata

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	GetAllPariwisata() (models.Respone, error)
	CreatePariwisata(pariwisataName, pariwisataLokasi string) error
}
