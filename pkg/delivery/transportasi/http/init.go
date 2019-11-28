package http

import (
	"github.com/labstack/echo"

	transportasiUsecase "github.com/TobaTourism/pkg/usecase/transportasi"
)

type transportasi struct {
	transportasiUsecase transportasiUsecase.Usecase
}

func InitTransportasiHandler(e *echo.Echo, p transportasiUsecase.Usecase) {
	handler := &transportasi{
		transportasiUsecase: p,
	}

	// Handle GET
	e.GET("api/transportasi/getAll", handler.GetAllTransportasi)
	e.GET("api/transportasi/:id", handler.GetTransportasiByID)

	// Handle POST
	e.POST("api/transportasi/create", handler.CreateTransportasi)

	// Handle PUT
	e.PUT("api/transportasi/update/:id", handler.UpdateTransportasi)

	// Handle DELETE
	e.DELETE("api/transportasi/delete/:id", handler.DeleteTransportasi)
}
