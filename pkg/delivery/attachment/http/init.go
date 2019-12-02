package http

import (
	"github.com/labstack/echo"

	attachmentUsecase "github.com/TobaTourism/pkg/usecase/attachment"
)

type attachment struct {
	attachmentUsecase attachmentUsecase.Usecase
}

func InitAttachmentHandler(e *echo.Echo, p attachmentUsecase.Usecase) {
	handler := &attachment{
		attachmentUsecase: p,
	}

	//register handler
	//e.GET("/attachment", handler.GetAttachment)
	e.GET("/file", handler.ShowAttachment)
}
