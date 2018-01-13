package usecase

import (
	"github.com/yaaaaashiki/cstack/domain/repository"
	"go/ast"
)

type SetCurrentObjectStatusUseCase struct {
	objectRepository *repository.ObjectRepository
}

func NewSetCurrentObjectStatusUseCase(objectRepository *repository.ObjectRepository) *SetCurrentObjectStatusUseCase {
	return &SetCurrentObjectStatusUseCase{
		objectRepository: objectRepository,
	}
}
