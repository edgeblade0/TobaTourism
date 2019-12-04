package resto

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	GetAllResto() ([]models.Restoran, []int64, error)
	CreateResto(resto models.Restoran) error
}
