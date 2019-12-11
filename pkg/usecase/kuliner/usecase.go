package kuliner

import "github.com/TobaTourism/pkg/models"

type Usecase interface {
	CreateKuliner(restoID string, name string, desc string, price string, attachmentID int64) error
	GetAllCulinary(restoID int64) ([]models.Culinary, error)
	GetDetailKuliner(restoID string, culinaryID string) (models.Culinary, error)
	UpdateKuliner(restoID string, kulinerID string, kulinerName string, kulinerDesc string, kulinerPrice string) error
	UpdateImageKuliner(restoID string, kulinerID string, attachmentID int64) error
	DeleteKuliner(restoID string, kulinerID string) error
}
