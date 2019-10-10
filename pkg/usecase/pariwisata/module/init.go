package module

import (
	pariwisataRepo "github.com/TobaTourism/pkg/repository/pariwisata"
	pariwisataUsecase "github.com/TobaTourism/pkg/usecase/pariwisata"
)

type pariwisata struct {
	pariwisataRepo pariwisataRepo.Repository
}

func InitPariwisataUsecase(p pariwisataRepo.Repository) pariwisataUsecase.Usecase {
	return &pariwisata{
		pariwisataRepo: p,
	}
}
