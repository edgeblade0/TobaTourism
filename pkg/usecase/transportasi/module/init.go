package module

import (
	attachmentRepo "github.com/TobaTourism/pkg/repository/attachment"
	transportasiRepo "github.com/TobaTourism/pkg/repository/transportasi"
	transportasiUsecase "github.com/TobaTourism/pkg/usecase/transportasi"
)

type transportasi struct {
	transportasiRepo transportasiRepo.Repository
	attachmentRepo   attachmentRepo.Repository
}

// InitTransportasiUsecase Initialize Transportasi Usecase
func InitTransportasiUsecase(p transportasiRepo.Repository, a attachmentRepo.Repository) transportasiUsecase.Usecase {
	return &transportasi{
		transportasiRepo: p,
		attachmentRepo:   a,
	}
}
