package kuliner

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	CreateKuliner(culinary models.Culinary) error
	GetAllKuliner(RestoID int64) ([]models.Culinary, []int64, error)
}
