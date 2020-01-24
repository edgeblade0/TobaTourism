package transportasi

import (
	"github.com/labstack/echo"
)

type Delivery interface {
	GetAllTransportasi(c echo.Context) error
	GetTransportasiByID(c echo.Context) error
	CreateTransportasi(c echo.Context) error
	UpdateTransportasi(c echo.Context) error
	UpdateImageTransportasi(c echo.Context) error
	DeleteTransportasi(c echo.Context) error
}
