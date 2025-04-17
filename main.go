package main

import (
	"fmt"
)

func dadosPessoa(nome string, idade int) (int, string) {
	var status string
	if idade >= 18 {
		status = "maior de idade"
	} else {
		status = "menor de idade"
	}
	return idade, status
}

func main() {
	nome := "Miguel"
	idade := 17

	idadeRetornada, status := dadosPessoa(nome, idade)

	fmt.Println("Nome: ", nome )
	fmt.Println("Idade: ", idadeRetornada)
	fmt.Println("Status: ", status )
}
