package kuliner

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	CreateKuliner(culinary models.Culinary) error
	GetAllKuliner(RestoID int64) ([]models.Culinary, []int64, error)
	GetDetailKuliner(RestoID int64, CulinaryID int64) (models.Culinary, int64, error)
	UpdateKuliner(kuliner models.Culinary) error
	UpdateImageKuliner(kuliner models.Culinary) error
	DeleteKuliner(kuliner models.Culinary) error
	DeleteAllKuliner(restoID int64) error
}
