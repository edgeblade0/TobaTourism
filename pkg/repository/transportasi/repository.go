package transportasi

import (
	"github.com/TobaTourism/pkg/models"
)

type Repository interface {
	GetAllTransportasi() (models.TransportasiResponse, []int64, error)
	GetTransportasiByID(transportasiID int64) (models.TransportasiDetailResponse, error)
	CreateTransportasi(nama, rute, description, contact string, harga int64, attachmentID int64) (models.TransportasiResponse, error)
	UpdateTransportasi(transportasiID int64, nama, rute, description, contact string, harga int64) (models.TransportasiResponse, error)
	UpdateImageTransportasi(transportasiID int64, attachmentID int64) (models.TransportasiResponse, error)
	DeleteTransportasi(transportasiID int64) (models.TransportasiResponse, error)
}
