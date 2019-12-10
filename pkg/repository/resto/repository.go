package resto

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	GetAllResto() ([]models.Restoran, []int64, error)
	CreateResto(resto models.Restoran) error
	UpdateImageResto(resto models.Restoran) error
	UpdateResto(resto models.Restoran) error
	DeleteResto(restoID int64) error
}
