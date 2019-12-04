package http

import (
	"github.com/labstack/echo"

	attachmentUsecase "github.com/TobaTourism/pkg/usecase/attachment"
	kulinerUsecase "github.com/TobaTourism/pkg/usecase/kuliner"
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
	// e.GET("/resto", handler.GetAllResto)
	e.POST("/api/kuliner/create/:restoranId", handler.CreateKuliner)
}
