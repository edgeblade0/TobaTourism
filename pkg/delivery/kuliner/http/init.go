package http

import (
	"github.com/labstack/echo"

	attachmentUsecase "github.com/if416021/TobaTourism/pkg/usecase/attachment"
	kulinerUsecase "github.com/if416021/TobaTourism/pkg/usecase/kuliner"
)

type kuliner struct {
	kulinerUsecase    kulinerUsecase.Usecase
	attachmentUsecase attachmentUsecase.Usecase
}

func InitKulinerHandler(e *echo.Echo, p kulinerUsecase.Usecase, a attachmentUsecase.Usecase) {
	handler := &kuliner{
		kulinerUsecase:    p,
		attachmentUsecase: a,
	}

	// handler
	e.POST("/api/restaurant/:restoranId/culinary", handler.CreateKuliner)
	e.GET("/api/restaurant/:restaurantId/culinary/:culinaryId", handler.GetDetailKuliner)
	e.PUT("/api/restaurant/:restaurantId/culinary/:culinaryId", handler.UpdateKuliner)
	e.PUT("/api/restaurant/image/:restaurantId/culinary/:culinaryId", handler.UpdateImageKuliner)
	e.DELETE("/api/restaurant/:restaurantId/culinary/:culinaryId", handler.DeleteKuliner)
}
