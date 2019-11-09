package Pariwisata

import (
	"github.com/labstack/echo"
)

type Delivery interface {
	GetAllPariwisata(c echo.Context) error
	CreatePariwisata(c echo.Context) error
}
