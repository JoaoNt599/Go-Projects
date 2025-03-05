package use_cases

import (
	"log"
	"ms-categories/internal/entities"
	"ms-categories/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{
		repository: repository,
	}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	// TODO: persist entity to db
	log.Println(category)

	// TODO: verify id category name already exists
	err = u.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}
