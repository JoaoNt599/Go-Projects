package repositories

import (
	"errors"
	"ms-categories/internal/entities"
	"strconv"
)

type inMemoryCategoryRepository struct {
	db []*entities.Category
}

func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *inMemoryCategoryRepository) Save(category *entities.Category) error {
	// add an ID berfore saving
	r.db = append(r.db, category)

	return nil
}

func (r *inMemoryCategoryRepository) List() ([]*entities.Category, error) {
	return r.db, nil
}

func (r *inMemoryCategoryRepository) Update(id string, name string) (*entities.Category, error) {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, errors.New("invalid ID")
	}

	for _, cat := range r.db {
		if cat.ID == uint(idUint) {
			cat.Name = name
			return cat, nil
		}
	}
	return nil, errors.New("category not found")
}

func (r *inMemoryCategoryRepository) Delete(id string) error {
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("invalid ID")
	}

	for i, cat := range r.db {
		if cat.ID == uint(idUint) {
			r.db = append(r.db[:i], r.db[i+1:]...)
			return nil
		}
	}

	return errors.New("category not found")
}
