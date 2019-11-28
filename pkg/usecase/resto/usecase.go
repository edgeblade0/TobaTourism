package resto

import "github.com/TobaTourism/pkg/models"

type Usecase interface {
	GetAllResto() ([]models.Restoran, error)
	CreateResto(name string, location string, contact string, attachmentID int64) error
}
