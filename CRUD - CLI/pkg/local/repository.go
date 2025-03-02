package local

import (
	"consumer/internal/domain"
	"encoding/json"
	"io/ioutil"
)

type LocalRepository struct {
	fileName string
}

func NewLocalRepository(fileName string) *LocalRepository {
	return &LocalRepository{fileName: fileName}
}

func (r *LocalRepository) Read() ([]domain.Item, error) {
	file, err := ioutil.ReadFile(r.fileName)
	if err != nil {
		return nil, err
	}

	var items []domain.Item
	json.Unmarshal(file, &items)
	return items, nil
}

func (r *LocalRepository) Write(items []domain.Item) error {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(r.fileName, data, 0644)
}
