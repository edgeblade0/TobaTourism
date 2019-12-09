package http

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"

	"github.com/TobaTourism/pkg/models"
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
	restoID := c.Param("id")
	name := c.FormValue("culinaryName")
	desc := c.FormValue("culinaryDescriptiion")
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
