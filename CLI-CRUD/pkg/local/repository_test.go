package local

import (
	"consumer/internal/domain"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteAndRead(t *testing.T) {
	fileName := "test_items.json"
	repo := NewLocalRepository(fileName)
	defer os.Remove(fileName)

	items := []domain.Item{
		{ID: 1, Name: "Monitor", Price: 700.0},
	}

	err := repo.Write(items)
	assert.Nil(t, err)

	readItems, err := repo.Read()
	assert.Nil(t, err)
	assert.Equal(t, len(items), len(readItems))
	assert.Equal(t, "Monitor", readItems[0].Name)
}
