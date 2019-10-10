package http

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (d *pariwisata) GetAllPariwisata(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	allPariwisata, err := d.pariwisataUsecase.GetAllPariwisata()
	if err != nil {
		log.Println(err)
	}

	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, allPariwisata)
}
