package kuliner

import (
	"github.com/labstack/echo"
)

type delivery interface {
	CreateKuliner(c echo.Context) error
	GetDetailKuliner(c echo.Context) error
	UpdateKuliner(c echo.Context) error
	UpdateImageKuliner(c echo.Context) error
	DeleteKuliner(c echo.Context) error
}
