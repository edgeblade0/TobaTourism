package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
)

// func (d *attachment) GetAttachment(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	if ctx == nil {
// 		ctx = context.Background()
// 	}

// 	attachmentID := c.FormValue("attachment_id")

// 	err := d.attachmentUsecase.GetAttachment(attachmentID)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	c.Response().Header().Set(`X-Cursor`, "header")
// 	return c.JSON(http.StatusOK, nil)
// }

func (d *attachment) ShowAttachment(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	// attachmentID := c.FormValue("attachment_id")

	// err := d.attachmentUsecase.GetAttachment(attachmentID)
	// if err != nil {
	// 	log.Println(err)
	// }
	pathFile := c.FormValue("path")
	c.File(pathFile)

	c.Response().Header().Set(`X-Cursor`, "header")
	return c.JSON(http.StatusOK, nil)
}
