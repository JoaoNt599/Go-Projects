package usecase

import "consumer/internal/domain"

type ItemService struct {
	Repo ItemRepository
}

func NewItemService(repo ItemRepository) *ItemService {
	return &ItemService{Repo: repo}
}

func (s *ItemService) Create(name string, price float64) error {
	items, _ := s.Repo.Read()
	id := len(items) + 1
	newItem := domain.Item{ID: id, Name: name, Price: price}
	items = append(items, newItem)
	return s.Repo.Write(items)
}

func (s *ItemService) List() ([]domain.Item, error) {
	return s.Repo.Read()
}
