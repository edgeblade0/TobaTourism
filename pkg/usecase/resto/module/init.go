package module

import (
	restoRepo "github.com/TobaTourism/pkg/repository/resto"
	restoUsecase "github.com/TobaTourism/pkg/usecase/resto"
	attachmentRepo "github.com/TobaTourism/pkg/repository/attachment"
	kulinerRepo "github.com/TobaTourism/pkg/repository/kuliner"
	kulinerUsecase "github.com/TobaTourism/pkg/usecase/kuliner"
)

type resto struct {
	restoRepo      restoRepo.Repository
	attachmentRepo attachmentRepo.Repository
	kulinerRepo    kulinerRepo.Repository
	kulinerUsecase kulinerUsecase.Usecase
}

func InitRestoUsecase(p restoRepo.Repository, a attachmentRepo.Repository, k kulinerRepo.Repository, u kulinerUsecase.Usecase) restoUsecase.Usecase {
	return &resto{
		restoRepo:      p,
		attachmentRepo: a,
		kulinerRepo:    k,
		kulinerUsecase: u,
	}
}
