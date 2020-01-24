package http

import (
	"github.com/labstack/echo"

	attachmentUsecase "github.com/TobaTourism/pkg/usecase/attachment"
	restoUsecase "github.com/TobaTourism/pkg/usecase/resto"
)

type resto struct {
	restoUsecase      restoUsecase.Usecase
	attachmentUsecase attachmentUsecase.Usecase
}

func InitRestoHandler(e *echo.Echo, p restoUsecase.Usecase, a attachmentUsecase.Usecase) {
	handler := &resto{
		restoUsecase:      p,
		attachmentUsecase: a,
	}

	// handler
	e.POST("/api/restaurant", handler.InsertResto)
	e.POST("/api/restaurant/image/:restaurantID", handler.UpdateImageResto)
	e.PUT("/api/restaurant/:restaurantID", handler.UpdateResto)
	e.DELETE("/api/restaurant/:restaurantID", handler.DeleteResto)

	e.GET("/api/culinary", handler.GetAllRestoWithKuliner)
	e.GET("/api/restaurant", handler.GetAllResto)
	e.GET("api/restaurant/:restaurantID", handler.GetDetailResto)
}
