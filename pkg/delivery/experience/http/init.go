package http

import (
	"github.com/labstack/echo"

	attachmentUsecase "github.com/TobaTourism/pkg/usecase/attachment"
	experienceUsecase "github.com/TobaTourism/pkg/usecase/experience"
)

type experience struct {
	experienceUsecase experienceUsecase.Usecase
	attachmentUsecase attachmentUsecase.Usecase
}

// InitExperienceHandler initialize handler
func InitExperienceHandler(e *echo.Echo, p experienceUsecase.Usecase, a attachmentUsecase.Usecase) {
	handler := &experience{
		experienceUsecase: p,
		attachmentUsecase: a,
	}

	// Handle GET
	e.GET("api/experience/getAll", handler.GetAllExperience)
	e.GET("api/experience/:id", handler.GetExperienceByID)

	// Handle POST
	e.POST("api/experience/create", handler.CreateExperience)

	// Handle PUT
	e.PUT("api/experience/update/:id", handler.UpdateExperience)

	// Handle DELETE
	e.DELETE("api/experience/delete/:id", handler.DeleteExperience)
}
