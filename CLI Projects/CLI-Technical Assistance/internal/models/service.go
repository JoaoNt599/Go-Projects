package models

import "fmt"

type Service struct {
	ID          int
	Description string
	Value       float64
}

func (s Service) Displayinformation() string {
	return "ID: " + string(s.ID) +
		", Description: " + s.Description +
		", Value: R$ " + fmt.Sprintf("%.2f", s.Value)
}
