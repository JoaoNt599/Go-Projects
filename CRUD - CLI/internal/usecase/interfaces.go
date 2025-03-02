package usecase

import "consumer/internal/domain"

type ItemRepository interface {
	Read() ([]domain.Item, error)
	Write([]domain.Item) error
}
