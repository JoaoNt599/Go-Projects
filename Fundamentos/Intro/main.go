package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func NewUser(nome string) Pessoa {

	p := Pessoa{
		Nome:  nome,
		Idade: 27,
	}

	return p
}

// FezAniversario recebe uma valor int e aualiza ele
func FezAniversario(i *int) {
	*i = *i + 1
}

func main() {

	var pessoas = []Pessoa{
		{"Joao", 27},
		{"Maria", 32},
		{"Ana", 18},
	}

	n := pessoas[0].Nome
	i := pessoas[0].Idade

	FezAniversario(&i)

	fmt.Println("Valor resgatado:", n, i)

	// for _, v := range pessoas {
	// 	fmt.Printf("Pessoa: %s\t\tidade %d\n", v.Nome, v.Idade)
	// }

	nome := "Jaum"
	pessoa := NewUser(nome)

	fmt.Printf("Pessoa: %s\t\tidade %d\n", pessoa.Nome, pessoa.Idade)
}
