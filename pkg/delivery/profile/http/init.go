package http

import (
	"github.com/labstack/echo"

	profileUsecase "github.com/TobaTourism/pkg/usecase/profile"
)

type profile struct {
	profileUsecase profileUsecase.Usecase
}

func InitProfileHandler(e *echo.Echo, p profileUsecase.Usecase) {
	handler := &profile{
		profileUsecase: p,
	}

	// handler
	e.GET("/api/profile/:userId", handler.GetProfile)
	e.PUT("/api/profile/:userId", handler.UpdateProfile)
}
