package http

import (
	"context"
	"log"
	"net/http"

	"github.com/TobaTourism/pkg/models"
	"github.com/labstack/echo"
)

func (d *resto) GetAllResto(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	data, err := d.restoUsecase.GetAllResto()
	if err != nil {
		log.Println(err)
	}
	resp.Data = data

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *resto) InsertResto(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	name := c.FormValue("restoran_name")
	location := c.FormValue("restoran_lokasi")
	contact := c.FormValue("restoran_contact")

	//multipart
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("[Delivery][Restoran][MultipartForm] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}
	files := form.File["image"]
	attachmentID, err := d.attachmentUsecase.InsertAttachment(files, models.PathFileRestoran, models.RestoranTypeAttachment)
	if err != nil {
		log.Println("[Delivery][Restoran][InsertAttachment] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	err = d.restoUsecase.CreateResto(name, location, contact, attachmentID)
	if err != nil {
		log.Println("[Delivery][Restoran][CreateResto] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}
