package domain

import "time"

type Users []struct {
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Age       string    `json:"Age"`
	ID        string    `json:"id"`
}
