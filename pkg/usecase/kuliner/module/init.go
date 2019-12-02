package module

import (
	kulinerRepo "github.com/TobaTourism/pkg/repository/kuliner"
	kulinerUsecase "github.com/TobaTourism/pkg/usecase/kuliner"
)

type kuliner struct {
	kulinerRepo kulinerRepo.Repository
}

func InitKulinerUsecase(p kulinerRepo.Repository) kulinerUsecase.Usecase {
	return &kuliner{
		kulinerRepo: p,
	}
}
