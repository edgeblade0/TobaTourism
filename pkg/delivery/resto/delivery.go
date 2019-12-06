package resto

import (
	"github.com/labstack/echo"
)

type delivery interface {
	GetAllRestoWithKuliner(c echo.Context) error
	GetDetailResto(c echo.Context) error
	InsertResto(c echo.Context) error
}
