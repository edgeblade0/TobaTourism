package resto

import (
	"github.com/labstack/echo"
)

type delivery interface {
	GetAllResto(c echo.Context) error
	InsertResto(c echo.Context) error
}
