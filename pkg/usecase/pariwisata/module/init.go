package module

import (
	attachmentRepo "github.com/TobaTourism/pkg/repository/attachment"
	pariwisataRepo "github.com/TobaTourism/pkg/repository/pariwisata"
	pariwisataUsecase "github.com/TobaTourism/pkg/usecase/pariwisata"
)

type pariwisata struct {
	attachmentRepo attachmentRepo.Repository
	pariwisataRepo pariwisataRepo.Repository
}

func InitPariwisataUsecase(p pariwisataRepo.Repository, a attachmentRepo.Repository) pariwisataUsecase.Usecase {
	return &pariwisata{
		pariwisataRepo: p,
		attachmentRepo: a,
	}
}
