package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/TobaTourism/pkg/models"
)

func (d *resto) GetAllRestoWithKuliner(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	data, err := d.restoUsecase.GetAllRestoWithKuliner()
	if err != nil {
		log.Println(err)
	}
	resp.Data = data

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

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

func (d *resto) GetDetailResto(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	restoID, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	data, err := d.restoUsecase.GetDetailResto(restoID)
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
	name := c.FormValue("restaurantName")
	location := c.FormValue("restaurantLocation")
	contact := c.FormValue("restaurantContact")

	//multipart
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("[Delivery][Restoran][MultipartForm] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}
	files := form.File["restaurantImage"]
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
