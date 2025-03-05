package main

import (
	"consumer/internal/usecase"
	"consumer/pkg/local"
	"fmt"
)

func main() {
	repo := local.NewLocalRepository("items.json")
	service := usecase.NewItemService(repo)

	for {
		fmt.Println("\n1. Criar Item")
		fmt.Println("2. Listar Itens")
		fmt.Println("3. Sair")
		fmt.Print("Escolha: ")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			var name string
			var price float64
			fmt.Print("Nome: ")
			fmt.Scanln(&name)
			fmt.Print("Preço: ")
			fmt.Scanln(&price)
			service.Create(name, price)
			fmt.Println("Item criado")
		case 2:
			items, _ := service.List()
			for _, item := range items {
				fmt.Printf("ID: %d, Nome: %s, Preço: %.2f\n", item.ID, item.Name, item.Price)
			}
		case 3:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
}
