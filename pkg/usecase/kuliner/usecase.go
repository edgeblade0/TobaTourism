package kuliner

import "github.com/TobaTourism/pkg/models"

type Usecase interface {
	//GetAllResto() ([]models.Restoran, error)
	CreateKuliner(restoID string, name string, desc string, price string, attachmentID int64) error
	GetAllCulinary(restoID int64) ([]models.Culinary, error)
}
