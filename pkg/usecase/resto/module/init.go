package module

import (
	restoRepo "github.com/TobaTourism/pkg/repository/resto"
	restoUsecase "github.com/TobaTourism/pkg/usecase/resto"
)

type resto struct {
	restoRepo restoRepo.Repository
}

func InitRestoUsecase(p restoRepo.Repository) restoUsecase.Usecase {
	return &resto{
		restoRepo: p,
	}
}
