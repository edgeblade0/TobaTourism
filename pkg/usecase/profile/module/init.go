package module

import (
	profileRepo "github.com/TobaTourism/pkg/repository/profile"
	profileUsecase "github.com/TobaTourism/pkg/usecase/profile"
)

type profile struct {
	profileRepo profileRepo.Repository
}

func InitProfileUsecase(p profileRepo.Repository) profileUsecase.Usecase {
	return &profile{
		profileRepo: p,
	}
}
