package experience

import (
	"github.com/labstack/echo"
)

// Delivery interface
type Delivery interface {
	GetAllExperience(c echo.Context) error
	GetExperienceByID(c echo.Context) error
	CreateExperience(c echo.Context) error
	UpdateExperience(c echo.Context) error
	DeleteExperience(c echo.Context) error
}
