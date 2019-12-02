package module

import (
	attachmentRepo "github.com/TobaTourism/pkg/repository/attachment"
	attachmentUsecase "github.com/TobaTourism/pkg/usecase/attachment"
)

type attachment struct {
	attachmentRepo attachmentRepo.Repository
}

func InitAttachmentUsecase(p attachmentRepo.Repository) attachmentUsecase.Usecase {
	return &attachment{
		attachmentRepo: p,
	}
}
