package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (d *transportasi) GetAllTransportasi(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	allTransportasi, err := d.transportasiUsecase.GetAllTransportasi()
	if err != nil {
		log.Println("[Delivery][Transportasi][GetAllTransportasi] Error: ", err)

		return c.JSON(http.StatusInternalServerError, allTransportasi)
	}

	c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, allTransportasi)
}

func (d *transportasi) GetTransportasiByID(c echo.Context) error {
	transportasiID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	transportasi, err := d.transportasiUsecase.GetTransportasiByID(transportasiID)
	if err != nil {
		log.Println("[Delivery][Transportasi][GetTransportasiByID] Error: ", err)

		return c.JSON(http.StatusInternalServerError, transportasi)
	}

	c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, transportasi)
}

func (d *transportasi) CreateTransportasi(c echo.Context) error {
	nama := c.FormValue("nama")
	rute := c.FormValue("rute")
	description := c.FormValue("description")
	contact := c.FormValue("contact")
	harga, _ := strconv.ParseInt(c.FormValue("harga"), 10, 64)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	transportasi, err := d.transportasiUsecase.CreateTransportasi(nama, rute, description, contact, harga)
	if err != nil {
		log.Println("[Delivery][Transportasi][CreateTransportasi] Error: ", err)

		return c.JSON(http.StatusInternalServerError, transportasi)
	}

	return c.JSON(http.StatusOK, transportasi)
}

func (d *transportasi) UpdateTransportasi(c echo.Context) error {
	transportasiID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	nama := c.FormValue("nama")
	rute := c.FormValue("rute")
	description := c.FormValue("description")
	contact := c.FormValue("contact")
	harga, _ := strconv.ParseInt(c.FormValue("harga"), 10, 64)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	transportasi, err := d.transportasiUsecase.UpdateTransportasi(transportasiID, nama, rute, description, contact, harga)
	if err != nil {
		log.Println("[Delivery][Transportasi][UpdateTransportasi] Error: ", err)

		return c.JSON(http.StatusInternalServerError, transportasi)
	}

	return c.JSON(http.StatusOK, transportasi)
}

func (d *transportasi) DeleteTransportasi(c echo.Context) error {
	transportasiID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	transportasi, err := d.transportasiUsecase.DeleteTransportasi(transportasiID)
	if err != nil {
		log.Println("[Delivery][Transportasi][DeleteTransportasi] Error: ", err)

		return c.JSON(http.StatusInternalServerError, transportasi)
	}

	return c.JSON(http.StatusOK, transportasi)
}
