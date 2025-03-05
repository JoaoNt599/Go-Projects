package models

type Technical struct {
	Name  string
	CPF   string
	Phone string
}

func (t Technical) Displayinformation() string {
	return "Name: " + t.Name +
		", CPF: " + t.CPF +
		", Phone: " + t.Phone
}
