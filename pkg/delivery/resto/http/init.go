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
	e.GET("/api/culinary", handler.GetAllResto)
	e.POST("/api/restoran/create", handler.InsertResto)
}
