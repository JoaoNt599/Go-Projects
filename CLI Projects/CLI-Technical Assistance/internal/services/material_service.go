package services

import (
	"fmt"
	"technical-assistence/internal/models"
)

func ListMaterials(materials []models.Material) {
	for _, m := range materials {
		fmt.Println(m.Displayinformation())
	}
}
