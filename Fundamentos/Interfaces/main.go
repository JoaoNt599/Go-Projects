package main

import "fmt"

type Humano interface {
	Fala(text string)
	Sente(sentimento string)
}

type Robo struct {
	Nome           string
	DataFabricacao string
}

type Pessoa struct {
	Nome  string
	Idade int
}

func (pessoa Pessoa) Fala(texto string) {
	fmt.Print(texto)
}

func (robo Robo) Fala(texto string) {
	fmt.Println(texto)
}

func (pessoa Pessoa) Sente(sentimento string) {
	fmt.Println("Estou sentidno", sentimento)
}

func CadastrarCPF(h Humano) {
	// h.CPF = "12345678910"
	fmt.Println("CPF Cadastrado")
}

func main() {
	pessoa := &Pessoa{}
	robo := &Robo{}

	robo.Fala("Hello")

	pessoa.Fala("Olá")

	CadastrarCPF(pessoa)
}
