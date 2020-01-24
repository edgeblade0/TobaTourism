package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/TobaTourism/pkg/models"
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

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	nama := c.FormValue("transportationName")
	rute := c.FormValue("transportationRoute")
	description := c.FormValue("transportationDescription")
	contact := c.FormValue("transportationContact")
	harga, _ := strconv.ParseInt(c.FormValue("transportationPrice"), 10, 64)

	//multipart
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("[Delivery][Transportasi][MultipartForm] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, nil)
	}
	files := form.File["transportationImage"]
	attachmentID, err := d.attachmentUsecase.InsertAttachment(files, models.PathFileTransportasi, models.TransportasiTypeAttachment)
	if err != nil {
		log.Println("[Delivery][Transportasi][InsertAttachment] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, nil)
	}
	transportasi, err := d.transportasiUsecase.CreateTransportasi(nama, rute, description, contact, harga, attachmentID)
	if err != nil {
		log.Println("[Delivery][Transportasi][CreateTransportasi] Error: ", err)

		return c.JSON(http.StatusInternalServerError, transportasi)
	}

	return c.JSON(http.StatusOK, transportasi)
}

func (d *transportasi) UpdateTransportasi(c echo.Context) error {
	transportasiID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	nama := c.FormValue("transportationName")
	rute := c.FormValue("transportationRoute")
	description := c.FormValue("transportationDescription")
	contact := c.FormValue("transportationContact")
	harga, _ := strconv.ParseInt(c.FormValue("transportationPrice"), 10, 64)

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

func (d *transportasi) UpdateImageTransportasi(c echo.Context) error {
	transportasiID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	//multipart
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("[Delivery][Transportasi][MultipartForm] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, nil)
	}
	files := form.File["transportationImage"]
	attachmentID, err := d.attachmentUsecase.InsertAttachment(files, models.PathFileTransportasi, models.TransportasiTypeAttachment)
	if err != nil {
		log.Println("[Delivery][Transportasi][InsertAttachment] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, nil)
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	transportasi, err := d.transportasiUsecase.UpdateImageTransportasi(transportasiID, attachmentID)
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
