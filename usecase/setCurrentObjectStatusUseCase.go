package usecase

import "github.com/yaaaaashiki/livething/domain/repository"

type SetCurrentObjectStatusUseCase struct {
	objectRepository *repository.ObjectRepository
}

func NewSetCurrentObjectStatusUseCase(objectRepository *repository.ObjectRepository) *SetCurrentObjectStatusUseCase {
	return &SetCurrentObjectStatusUseCase{
		objectRepository: objectRepository,
	}
}
