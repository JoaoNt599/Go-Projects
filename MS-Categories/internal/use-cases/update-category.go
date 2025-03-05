package use_cases

import (
	"errors"
	"ms-categories/internal/repositories"
)

type updateCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewUpdateCategoryUseCase(repository repositories.ICategoryRepository) *updateCategoryUseCase {
	return &updateCategoryUseCase{
		repository: repository,
	}
}

func (u *updateCategoryUseCase) Execute(id string, name string) error {
	if name == "" {
		return errors.New("name is required")
	}

	_, err := u.repository.Update(id, name)
	if err != nil {
		return err
	}

	return nil
}
