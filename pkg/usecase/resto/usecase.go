package resto

import "github.com/TobaTourism/pkg/models"

type Usecase interface {
	GetAllResto() ([]models.Restoran, error)
	CreateResto(name string, location string, contact string, attachmentID int64) error
	UpdateImageRestoran(restoID string, attachmentID int64) error
	UpdateRestoran(restoID string, restoName string, restoContact string, restoLocation string) error
	DeleteResto(restoID string) error
}
