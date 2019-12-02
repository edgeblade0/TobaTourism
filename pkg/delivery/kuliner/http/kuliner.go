package http

import (
	"context"
	"log"
	"net/http"

	"github.com/TobaTourism/pkg/models"
	"github.com/labstack/echo"
)

// func (d *resto) GetAllResto(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	if ctx == nil {
// 		ctx = context.Background()
// 	}

// 	allResto, err := d.restoUsecase.GetAllResto()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	c.Response().Header().Set(`X-Cursor`, "header")
// 	return c.JSON(http.StatusOK, allResto)
// }

func (d *kuliner) CreateKuliner(c echo.Context) error {
	var resp models.Responses
	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	restoID := c.Param("restoranId")
	name := c.FormValue("kuliner_nama")
	desc := c.FormValue("kuliner_desc")
	price := c.FormValue("kuliner_harga")

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
