package http

import (
	"github.com/labstack/echo"

	attachmentUsecase "github.com/TobaTourism/pkg/usecase/attachment"
	transportasiUsecase "github.com/TobaTourism/pkg/usecase/transportasi"
)

type transportasi struct {
	transportasiUsecase transportasiUsecase.Usecase
	attachmentUsecase   attachmentUsecase.Usecase
}

func InitTransportasiHandler(e *echo.Echo, p transportasiUsecase.Usecase, a attachmentUsecase.Usecase) {
	handler := &transportasi{
		transportasiUsecase: p,
		attachmentUsecase:   a,
	}

	// Handle GET
	e.GET("api/transportasi/getAll", handler.GetAllTransportasi)
	e.GET("api/transportasi/:id", handler.GetTransportasiByID)

	// Handle POST
	e.POST("api/transportasi/create", handler.CreateTransportasi)

	// Handle PUT
	e.PUT("api/transportasi/:id", handler.UpdateTransportasi)
	e.POST("api/transportasi/image/:id", handler.UpdateImageTransportasi)

	// Handle DELETE
	e.DELETE("api/transportasi/:id", handler.DeleteTransportasi)
}
