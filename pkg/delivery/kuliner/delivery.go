package kuliner

import (
	"github.com/labstack/echo"
)

type delivery interface {
	//GetAllResto(c echo.Context) error
	CreateKuliner(c echo.Context) error
}
