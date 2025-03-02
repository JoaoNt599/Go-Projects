package main

import "fmt"

// Esse programa registra na memoria o site e o criador

type Site struct {
	Criador string
	URL     string
}

func (s *Site) DefineCriador(Name string) {
	s.Criador = Name
	fmt.Println("Dado salvo")
}

func (s *Site) DefineURL(URL string) {
	s.URL = URL
	fmt.Println("Dado savlo")
}

func main() {

	site := Site{}

	site.DefineCriador("João")
	site.DefineURL("https://meusite.com")
	fmt.Printf("Criador: %s\nSite: %s\n", site.Criador, site.URL)
}
