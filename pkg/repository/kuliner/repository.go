package kuliner

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	//GetAllResto() ([]models.Restoran, error)
	CreateKuliner(culinary models.Culinary) error
}
