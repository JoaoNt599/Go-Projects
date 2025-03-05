package models

type Attendant struct {
	Name string
	CPF  string
}

func (a Attendant) Displayinformation() string {
	return "Name: " + a.Name + ", CPF: " + a.CPF
}
