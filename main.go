package main

import (
	"fmt"
)

func main() {
	var vetor [5]int
	var soma int

	for i := 0; i < 5; i++ {
		fmt.Printf("Digite o %dº número inteiro: ", i+1)
		fmt.Scan(&vetor[i])
		soma += vetor[i]
	}

	fmt.Printf("A soma dos números digitados é: %d\n", soma)
}
