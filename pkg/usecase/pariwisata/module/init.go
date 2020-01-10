package module

import (
	pariwisataRepo "github.com/if416021/TobaTourism/pkg/repository/pariwisata"
	pariwisataUsecase "github.com/if416021/TobaTourism/pkg/usecase/pariwisata"
)

type pariwisata struct {
	pariwisataRepo pariwisataRepo.Repository
}

func InitPariwisataUsecase(p pariwisataRepo.Repository) pariwisataUsecase.Usecase {
	return &pariwisata{
		pariwisataRepo: p,
	}
}
