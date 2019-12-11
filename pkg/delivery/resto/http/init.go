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
	e.POST("/api/restoran/create", handler.InsertResto)
	e.PUT("/api/restauran/image/:restauranID", handler.UpdateImageResto)
	e.PUT("/api/restauran/:restauranID", handler.UpdateResto)
	e.DELETE("/api/restauran/:restauranID", handler.DeleteResto)

	e.GET("/api/culinary", handler.GetAllRestoWithKuliner)
	e.GET("/api/restaurant", handler.GetAllResto)
	e.GET("api/restaurant/:id", handler.GetDetailResto)
}
