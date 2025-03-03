package main

import (
	"consumer/pkg/web"
	"fmt"
)

func main() {
	fmt.Println("Request")

	users, err := web.GetUsers()
	if err != nil {
		fmt.Println("Erro: ", err)
	}

	for indice, valor := range users {
		fmt.Printf("[%d] User: %s\n", indice, valor.Name)
	}
}
