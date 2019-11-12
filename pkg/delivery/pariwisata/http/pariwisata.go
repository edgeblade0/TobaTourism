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
		log.Println("[Delivery][Pariwisata][GetAllPariwisata] Error : ", err)

		return c.JSON(http.StatusInternalServerError, allPariwisata)
	}

	c.Response().Header().Set(`Content-Type`, "application/json")
	return c.JSON(http.StatusOK, allPariwisata)
}

//for Insert wisata
func (d *pariwisata) CreatePariwisata(c echo.Context) error {

	// pName := c.QueryParam("name")
	// pLokasi := c.QueryParam("lokasi")

	pName := c.FormValue("name")
	pLokasi := c.FormValue("lokasi")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err := d.pariwisataUsecase.CreatePariwisata(pName, pLokasi)
	if err != nil {
		log.Println(err)
	}

	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, err)

}
