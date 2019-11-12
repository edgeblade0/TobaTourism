package pariwisata

import "github.com/TobaTourism/pkg/models"

type Usecase interface {
	GetAllPariwisata() (models.Respone, error)
	CreatePariwisata(pName, pLokasi string) error
}
