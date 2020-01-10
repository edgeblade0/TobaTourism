package http

import (
	"github.com/labstack/echo"

	pariwisataUsecase "github.com/if416021/TobaTourism/pkg/usecase/pariwisata"
)

type pariwisata struct {
	pariwisataUsecase pariwisataUsecase.Usecase
}

func InitPariwisataHandler(e *echo.Echo, p pariwisataUsecase.Usecase) {
	handler := &pariwisata{
		pariwisataUsecase: p,
	}

	//Handle GET
	e.GET("api/pariwisata/getAll", handler.GetAllPariwisata)
	e.GET("api/pariwisata/:id", handler.GetPariwisataByID)

	// Handle POST
	e.POST("api/pariwisata/create", handler.CreatePariwisata)

	// Handle PUT
	e.PUT("api/pariwisata/update/:id", handler.UpdatePariwisata)

	// Handle DELETE
	e.DELETE("api/pariwisata/delete/:id", handler.DeletePariwisata)
}
