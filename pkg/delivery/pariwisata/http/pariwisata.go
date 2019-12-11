package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/TobaTourism/pkg/models"
)

func (d *pariwisata) GetAllPariwisata(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	allPariwisata, err := d.pariwisataUsecase.GetAllPariwisata()
	if err != nil {
		log.Println("[Delivery][Pariwisata][GetAllPariwisata] Error: ", err)

		return c.JSON(http.StatusInternalServerError, allPariwisata)
	}

	c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, allPariwisata)
}

func (d *pariwisata) GetPariwisataByID(c echo.Context) error {
	pariwisataID, _ := strconv.ParseInt(c.Param("pariwisata_id"), 10, 64)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	pariwisata, err := d.pariwisataUsecase.GetPariwisataByID(pariwisataID)
	if err != nil {
		log.Println("[Delivery][Pariwisata][GetPariwisataByID] Error: ", err)

		return c.JSON(http.StatusInternalServerError, pariwisata)
	}

	c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, pariwisata)
}

func (d *pariwisata) CreatePariwisata(c echo.Context) error {
	nama := c.FormValue("nama")
	lokasi := c.FormValue("lokasi")
	description := c.FormValue("description")
	contact := c.FormValue("contact")

		form, err := c.MultipartForm()
	if err != nil {
		log.Println("[Delivery][Pariwisata][MultipartForm] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, err)
	}
	files := form.File["pariwsataImage"]
	attachmentID, err := d.attachmentUsecase.InsertAttachment(files, models.PathFilePariwisata, models.PariwisataTypeAttachment)
	if err != nil {
		log.Println("[Delivery][Pariwisata][InsertAttachment] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, err)
		}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	pariwisata, err := d.pariwisataUsecase.CreatePariwisata(nama, lokasi, description, contact, attachmentID)
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
		log.Println("[Delivery][Pariwisata][UpdatePariwisata] Error: ", err)

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
