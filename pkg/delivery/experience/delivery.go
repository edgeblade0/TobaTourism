package experience

import (
	"github.com/labstack/echo"
)

// Delivery interface
type Delivery interface {
	GetAllExperience(c echo.Context) error
}
