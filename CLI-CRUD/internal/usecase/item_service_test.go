package usecase

import (
	"consumer/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockRepository struct {
	items []domain.Item
}

func (m *MockRepository) Read() ([]domain.Item, error) {
	return m.items, nil
}

func (m *MockRepository) Write(items []domain.Item) error {
	m.items = items
	return nil
}

func TestCreate(t *testing.T) {
	repo := &MockRepository{}
	service := NewItemService(repo)

	err := service.Create("Garrafa", 45.5)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(repo.items))
	assert.Equal(t, "Garrafa", repo.items[0].Name)
}

func TestList(t *testing.T) {
	repo := &MockRepository{
		items: []domain.Item{
			{ID: 1, Name: "Whey", Price: 150.99},
		},
	}
	service := NewItemService(repo)

	intems, err := service.List()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(intems))
	assert.Equal(t, "Whey", intems[0].Name)
}
