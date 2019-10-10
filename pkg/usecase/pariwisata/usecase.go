package pariwisata

import "github.com/TobaTourism/pkg/models"

type Usecase interface {
	GetAllPariwisata() ([]models.Pariwisata, error)
}
