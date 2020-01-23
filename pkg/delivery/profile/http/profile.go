package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/TobaTourism/pkg/models"
)

func (d *profile) GetProfile(c echo.Context) error {
	var resp models.Responses

	resp.Status = models.StatusFailed
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	userID := c.Param("userId")
	userIDInt64, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		log.Println("[Delivery][Profile][Parse userId Get] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	data, err := d.profileUsecase.GetProfile(userIDInt64)
	if err != nil {
		log.Println("[Delivery][Profile][GetProfile] Error : ", err)
		c.Response().Header().Set(`X-Cursor`, "header")
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Data = data
	resp.Message = models.MessageSucces
	if data.Name == "" {
		resp.Message = "Data tidak ditemukan"
	}
	resp.Status = models.StatusSucces
	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, resp)
}

func (d *profile) UpdateProfile(c echo.Context) error {

	return nil
}
