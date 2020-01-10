package http

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/if416021/TobaTourism/pkg/models"
)

func (d *kuliner) CreateKuliner(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	restoID := c.Param("restoranId")
	name := c.FormValue("culinaryName")
	desc := c.FormValue("culinaryDescription")
	price := c.FormValue("culinaryPrice")

	//multipart
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("[Delivery][Restoran][MultipartForm] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}
	files := form.File["culinaryImage"]
	attachmentID, err := d.attachmentUsecase.InsertAttachment(files, models.PathFileKuliner, models.KulinerTypeAttachment)
	if err != nil {
		log.Println("[Delivery][Kuliner][InsertAttachment] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	err = d.kulinerUsecase.CreateKuliner(restoID, name, desc, price, attachmentID)
	if err != nil {
		log.Println("[Delivery][Restoran][CreateResto] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *kuliner) GetDetailKuliner(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	restauranID := c.Param("restaurantId")
	kulinerID := c.Param("culinaryId")

	data, err := d.kulinerUsecase.GetDetailKuliner(restauranID, kulinerID)
	if err != nil {
		log.Println("[Delivery][Kuliner][GetDetailKuliner] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}
	resp.Data = data

	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *kuliner) UpdateKuliner(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	restoID := c.Param("restaurantId")
	kulinerID := c.Param("culinaryId")
	kulienrName := c.FormValue("culinaryName")
	KulinerDesc := c.FormValue("culinaryDescription")
	kulinerPrice := c.FormValue("culinaryPrice")

	err := d.kulinerUsecase.UpdateKuliner(restoID, kulinerID, kulienrName, KulinerDesc, kulinerPrice)
	if err != nil {
		log.Println("[Delivery][Restoran][UpdateKuliner] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *kuliner) UpdateImageKuliner(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	restoID := c.Param("restaurantId")
	kulinerID := c.Param("culinaryId")

	//multipart
	form, err := c.MultipartForm()
	if err != nil {
		log.Println("[Delivery][Restoran][MultipartForm] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}
	files := form.File["image"]
	attachmentID, err := d.attachmentUsecase.InsertAttachment(files, models.PathFileKuliner, models.KulinerTypeAttachment)
	if err != nil {
		log.Println("[Delivery][Kuliner][InsertAttachment on Update] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	err = d.kulinerUsecase.UpdateImageKuliner(restoID, kulinerID, attachmentID)
	if err != nil {
		log.Println("[Delivery][Restoran][Update] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *kuliner) DeleteKuliner(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	restoID := c.Param("restaurantId")
	kulinerID := c.Param("culinaryId")

	err := d.kulinerUsecase.DeleteKuliner(restoID, kulinerID)
	if err != nil {
		log.Println("[Delivery][Restoran][Delete] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Status = models.StatusSucces
	resp.Message = models.MessageSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}
