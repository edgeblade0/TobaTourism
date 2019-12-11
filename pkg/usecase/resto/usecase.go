package resto

import "github.com/TobaTourism/pkg/models"

type Usecase interface {
	GetAllRestoWithKuliner() ([]models.Restoran, error)
	GetAllResto() ([]models.Restoran, error)
	GetDetailResto(restoID int64) (models.Restoran, error)
	CreateResto(name string, location string, contact string, attachmentID int64) error
	UpdateImageRestoran(restoID string, attachmentID int64) error
	UpdateRestoran(restoID string, restoName string, restoContact string, restoLocation string) error
	DeleteResto(restoID string) error
}
