package module

import (
	kulinerRepo "github.com/TobaTourism/pkg/repository/kuliner"
	kulinerUsecase "github.com/TobaTourism/pkg/usecase/kuliner"

	attachmentRepo "github.com/TobaTourism/pkg/repository/attachment"
)

type kuliner struct {
	kulinerRepo    kulinerRepo.Repository
	attachmentRepo attachmentRepo.Repository
}

func InitKulinerUsecase(p kulinerRepo.Repository, a attachmentRepo.Repository) kulinerUsecase.Usecase {
	return &kuliner{
		kulinerRepo:    p,
		attachmentRepo: a,
	}
}
