package profile

import (
	"github.com/labstack/echo"
)

type delivery interface {
	GetProfile(c echo.Context) error
	UpdateProfile(c echo.Context) error
}
