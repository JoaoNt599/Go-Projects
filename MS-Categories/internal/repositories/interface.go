package repositories

import "ms-categories/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
	Delete(id string) error
	Update(id string, name string) (*entities.Category, error)
}
