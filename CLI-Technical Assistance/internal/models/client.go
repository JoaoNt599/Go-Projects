package models

type Client struct {
	Name    string
	CPF     string
	Address string
	Phone   string
}

func (c Client) Displayinformation() string {
	return "Name: " + c.Name +
		", CPF: " + c.CPF +
		", Address: " + c.Address +
		", Phone: " + c.Phone
}
