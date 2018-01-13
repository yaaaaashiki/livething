package usecase

type SetCurrentObjectStatusUseCase struct {
	objectRepository *repository.ObjectRepository
}

func NewSetCurrentObjectStatusUseCase(objectRepository *repository.ObjectRepository) *SetCurrentObjectStatusUseCase {
	return &SetCurrentObjectStatusUseCase{
		objectRepository: objectRepository,
	}
}
