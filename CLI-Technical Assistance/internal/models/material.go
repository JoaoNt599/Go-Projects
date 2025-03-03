package models

import "fmt"

type Material struct {
	ID          int
	Description string
	Value       float64
}

func (m Material) Displayinformation() string {
	return "ID: " + string(m.ID) +
		", Description: " + m.Description +
		", Value: R$: " + fmt.Sprintf("%.2f", m.Value)
}
