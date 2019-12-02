package attachment

import (
	"github.com/labstack/echo"
)

type Delivery interface {
	GetAttachment(c echo.Context) error
}
