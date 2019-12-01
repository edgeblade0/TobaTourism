package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

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

func (d *pariwisata) GetPariwisataByID(c echo.Context) error {
	log.Println("Test")
	return nil
}

//for Insert wisata
func (d *pariwisata) CreatePariwisata(c echo.Context) error {
	nama := c.Param("nama")
	lokasi := c.FormValue("lokasi")
	description := c.FormValue("description")
	contact := c.FormValue("contact")
	log.Println(nama)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	pariwisata, err := d.pariwisataUsecase.CreatePariwisata(nama, lokasi, description, contact)
	if err != nil {
		log.Println("[Delivery][Pariwisata][CreatePariwisata] Error: ", err)

		return c.JSON(http.StatusInternalServerError, pariwisata)
	}

	return c.JSON(http.StatusOK, pariwisata)
}

func (d *pariwisata) UpdatePariwisata(c echo.Context) error {
	pariwisataID, _ := strconv.ParseInt(c.Param("pariwisata_id"), 10, 64)

	nama := c.FormValue("nama")
	lokasi := c.FormValue("lokasi")
	description := c.FormValue("description")
	contact := c.FormValue("contact")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	pariwisata, err := d.pariwisataUsecase.UpdatePariwisata(pariwisataID, nama, lokasi, description, contact)
	if err != nil {
		log.Println("[Delivery][Pariwisata][Updatepariwisata] Error: ", err)

		return c.JSON(http.StatusInternalServerError, pariwisata)
	}

	return c.JSON(http.StatusOK, pariwisata)
}

func (d *pariwisata) DeletePariwisata(c echo.Context) error {
	pariwisataID, _ := strconv.ParseInt(c.Param("pariwisata_id"), 10, 64)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	pariwisata, err := d.pariwisataUsecase.DeletePariwisata(pariwisataID)
	if err != nil {
		log.Println("[Delivery][Pariwisata][DeletePariwisata] Error: ", err)

		return c.JSON(http.StatusInternalServerError, pariwisata)
	}

	return c.JSON(http.StatusOK, pariwisata)
}
