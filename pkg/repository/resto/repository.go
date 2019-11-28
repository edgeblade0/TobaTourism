package resto

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	GetAllResto() ([]models.Restoran, error)
	CreateResto(resto models.Restoran) error
}
