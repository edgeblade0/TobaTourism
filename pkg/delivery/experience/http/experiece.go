package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	"github.com/TobaTourism/pkg/models"
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

func (d *experience) GetExperienceByID(c echo.Context) error {
	id := c.Param("id")
	experienceID, _ := strconv.ParseInt(id, 10, 64)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	experience, err := d.experienceUsecase.GetExperienceByID(experienceID)
	if err != nil {
		log.Println("[Delivery][Experience][GetExperienceByID] Error: ", err)

		return c.JSON(http.StatusInternalServerError, experience)
	}

	c.Response().Header().Set("Content-Type", "application/json")

	return c.JSON(http.StatusOK, experience)
}

func (d *experience) CreateExperience(c echo.Context) error {
	var resp models.ExperienceResponse

	description := c.FormValue("description")
	location := c.FormValue("location")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	fileInput, err := c.MultipartForm()
	if err != nil {
		log.Println("[Delivery][Experience][CreateExperience][MultipartForm] Error: ", err)

		return c.JSON(http.StatusInternalServerError, resp)
	}

	files := fileInput.File["experienceImage"]
	attachmentID, err := d.attachmentUsecase.InsertAttachment(files, models.PathFileExperience, models.ExperienceTypeAttachment)
	if err != nil {
		log.Println("[Delivery][Experience][CreateExperience][InsertAttachment] Error: ", err)

		return c.JSON(http.StatusInternalServerError, resp)
	}

	experience, err := d.experienceUsecase.CreateExperience(description, location, attachmentID)
	if err != nil {
		log.Println("[Delivery][Experience][CreateExperience] Error: ", err)

		return c.JSON(http.StatusInternalServerError, experience)
	}

	return c.JSON(http.StatusOK, experience)
}

func (d *experience) UpdateExperience(c echo.Context) error {
	var resp models.ExperienceResponse

	id := c.Param("id")
	experienceID, _ := strconv.ParseInt(id, 10, 64)

	description := c.FormValue("description")
	location := c.FormValue("location")

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	fileInput, err := c.MultipartForm()
	if err != nil {
		log.Println("[Delivery][Experience][UpdateExperience][MultipartForm] Error: ", err)

		return c.JSON(http.StatusInternalServerError, resp)
	}

	files := fileInput.File["experienceImage"]
	attachmentID, err := d.attachmentUsecase.InsertAttachment(files, models.PathFileExperience, models.ExperienceTypeAttachment)
	if err != nil {
		log.Println("[Delivery][Experience][UpdateExperience][InsertAttachment] Error: ", err)

		return c.JSON(http.StatusInternalServerError, resp)
	}

	experience, err := d.experienceUsecase.UpdateExperience(experienceID, description, location, attachmentID)
	if err != nil {
		log.Println("[Delivery][Experience][UpdateExperience] Error: ", err)

		return c.JSON(http.StatusInternalServerError, experience)
	}

	return c.JSON(http.StatusOK, experience)
}

func (d *experience) DeleteExperience(c echo.Context) error {
	id := c.Param("id")
	experienceID, _ := strconv.ParseInt(id, 10, 64)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	experience, err := d.experienceUsecase.DeleteExperience(experienceID)
	if err != nil {
		log.Println("[Delivery][Experience][DeleteExperience] Error: ", err)

		return c.JSON(http.StatusInternalServerError, experience)
	}

	return c.JSON(http.StatusOK, experience)
}
