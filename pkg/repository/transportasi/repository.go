package transportasi

import (
	"github.com/if416021/TobaTourism/pkg/models"
)

type Repository interface {
	GetAllTransportasi() (models.TransportasiResponse, error)
	GetTransportasiByID(transportasiID int64) (models.TransportasiResponse, error)
	CreateTransportasi(nama, rute, description, contact string, harga int64) (models.TransportasiResponse, error)
	// UpdateImageTransportasi(transportasi models.TransportasiResponse)
	UpdateTransportasi(transportasiID int64, nama, rute, description, contact string, harga int64) (models.TransportasiResponse, error)
	DeleteTransportasi(transportasiID int64) (models.TransportasiResponse, error)
}
