package http

import (
	"github.com/labstack/echo"

	attachmentUsecase "github.com/TobaTourism/pkg/usecase/attachment"
	pariwisataUsecase "github.com/TobaTourism/pkg/usecase/pariwisata"
)

type pariwisata struct {
	attachmentUsecase attachmentUsecase.Usecase
	pariwisataUsecase pariwisataUsecase.Usecase
}

func InitPariwisataHandler(e *echo.Echo, p pariwisataUsecase.Usecase, a attachmentUsecase.Usecase) {
	handler := &pariwisata{
		pariwisataUsecase: p,
		attachmentUsecase: a,
	}

	//Handle GET
	e.GET("api/pariwisata/getAll", handler.GetAllPariwisata)
	e.GET("api/pariwisata/:pariwisata_id", handler.GetPariwisataByID)

	// Handle POST
	e.POST("api/pariwisata/create", handler.CreatePariwisata)

	// Handle PUT
	e.PUT("api/pariwisata/update/:pariwisata_id", handler.UpdatePariwisata)
	e.PUT("/api/pariwisata/image/:pariwisata_id", handler.UpdateImagePariwisata)

	// Handle DELETE
	e.DELETE("api/pariwisata/delete/:pariwisata_id", handler.DeletePariwisata)

}
