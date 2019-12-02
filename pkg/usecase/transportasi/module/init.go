package module

import (
	transportasiRepo "github.com/TobaTourism/pkg/repository/transportasi"
	transportasiUsecase "github.com/TobaTourism/pkg/usecase/transportasi"
)

type transportasi struct {
	transportasiRepo transportasiRepo.Repository
}

// InitTransportasiUsecase Initialize Transportasi Usecase
func InitTransportasiUsecase(p transportasiRepo.Repository) transportasiUsecase.Usecase {
	return &transportasi{
		transportasiRepo: p,
	}
}
