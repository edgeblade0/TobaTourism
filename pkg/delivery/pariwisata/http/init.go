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
	e.GET("api/tourism", handler.GetAllPariwisata)
	e.GET("api/tourism/:tourismId", handler.GetPariwisataByID)

	// Handle POST
	e.POST("api/tourism", handler.CreatePariwisata)

	// Handle PUT
	e.PUT("api/tourism/:tourismId", handler.UpdatePariwisata)
	e.PUT("/api/tourism/image/:tourismId", handler.UpdateImagePariwisata) //not clear

	// Handle DELETE
	e.DELETE("api/tourism/:tourismId", handler.DeletePariwisata)

}
