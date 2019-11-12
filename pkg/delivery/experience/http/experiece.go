package http

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (d *experience) GetAllExperience(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	allExperience, err := d.experienceUsecase.GetAllExperience()
	if err != nil {
		log.Println("[Delivery][Experience][GetAllExperience] Error: ", err)

		return c.JSON(http.StatusInternalServerError, allExperience)
	}

	c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, allExperience)
}
