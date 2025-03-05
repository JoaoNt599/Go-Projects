package web

import (
	"consumer/internal/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUsers() (domain.Users, error) {

	// resp, err := http.Get("http://example.com/")
	resp, err := http.Get("https://646a8f077d3c1cae4ce2a7fc.mockapi.io/api/v1/users")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Body: %s\n", body)

	var users domain.Users
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("Erro ao converter")
		return nil, err
	}

	return users, nil
}
