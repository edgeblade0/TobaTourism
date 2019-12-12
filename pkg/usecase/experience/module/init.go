package module

import (
	attachmentRepo "github.com/TobaTourism/pkg/repository/attachment"
	experienceRepo "github.com/TobaTourism/pkg/repository/experience"
	experienceUsecase "github.com/TobaTourism/pkg/usecase/experience"
)

type experience struct {
	experienceRepo experienceRepo.Repository
	attachmentRepo attachmentRepo.Repository
}

// InitExperienceUsecase Initialize Experience Usecase
func InitExperienceUsecase(p experienceRepo.Repository, a attachmentRepo.Repository) experienceUsecase.Usecase {
	return &experience{
		experienceRepo: p,
		attachmentRepo: a,
	}
}
