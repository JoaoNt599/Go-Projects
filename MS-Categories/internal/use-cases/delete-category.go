package use_cases

import "ms-categories/internal/repositories"

type deleteCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewDeleteCategoryUseCase(repository repositories.ICategoryRepository) *deleteCategoryUseCase {
	return &deleteCategoryUseCase{
		repository: repository,
	}
}

func (u *deleteCategoryUseCase) Execute(id string) error {
	return u.repository.Delete(id)
}
