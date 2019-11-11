package module

import (
	experienceRepo "github.com/TobaTourism/pkg/repository/experience"
	experienceUsecase "github.com/TobaTourism/pkg/usecase/experience"
)

type experience struct {
	experienceRepo experienceRepo.Repository
}

// InitExperienceUsecase Initialize Experience Usecase
func InitExperienceUsecase(p experienceRepo.Repository) experienceUsecase.Usecase {
	return &experience{
		experienceRepo: p,
	}
}
