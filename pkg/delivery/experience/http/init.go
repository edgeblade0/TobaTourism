package http

import (
	"github.com/labstack/echo"

	experienceUsecase "github.com/TobaTourism/pkg/usecase/experience"
)

type experience struct {
	experienceUsecase experienceUsecase.Usecase
}

// InitExperienceHandler initialize handler
func InitExperienceHandler(e *echo.Echo, p experienceUsecase.Usecase) {
	handler := &experience{
		experienceUsecase: p,
	}

	e.GET("/experience", handler.GetAllExperience)
}
