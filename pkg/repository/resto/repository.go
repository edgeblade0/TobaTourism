package resto

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	// GetAllRestoWithKuliner() ([]models.Restoran, []int64, error)
	GetAllResto() ([]models.Restoran, []int64, error)
	GetDetailResto(restoId int64) (models.Restoran, error)
	CreateResto(resto models.Restoran) error
	UpdateResto(resto models.Restoran) error
}
