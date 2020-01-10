package module

import (
	attachmentRepo "github.com/if416021/TobaTourism/pkg/repository/attachment"
	attachmentUsecase "github.com/if416021/TobaTourism/pkg/usecase/attachment"
)

type attachment struct {
	attachmentRepo attachmentRepo.Repository
}

func InitAttachmentUsecase(p attachmentRepo.Repository) attachmentUsecase.Usecase {
	return &attachment{
		attachmentRepo: p,
	}
}
