package Pariwisata

import (
	"github.com/labstack/echo"
)

type Delivery interface {
	GetAllPariwisata(c echo.Context) error
	GetPariwisataByID(c echo.Context) error
	CreatePariwisata(c echo.Context) error
	UpdatePariwisata(c echo.Context) error
	DeletePariwisata(c echo.Context) error
	UpdateImagePariwisata(c echo.Context) error
}
