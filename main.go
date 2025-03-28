package main 

import "fmt" 

func main() {
	var num1, num2 float32
	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)
	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	soma := num1 + num2
	subtracao := num1 - num2
	multiplicacao := num1 * num2
	divisao := num1 / num2

	fmt.Println("Resultados:")
	fmt.Printf("Número 1: %.2f\n", num1)
	fmt.Printf("Número 2: %.2f\n", num2)
	fmt.Printf("Soma: %.2f + %.2f = %.2f\n", num1, num2, soma)
	fmt.Printf("Subtração: %.2f - %.2f = %.2f\n", num1, num2, subtracao)
	fmt.Printf("Multiplicação: %.2f * %.2f = %.2f\n", num1, num2, multiplicacao)
	
	
	if num2 != 0 {
		fmt.Printf("Divisão: %.2f / %.2f = %.2f\n", num1, num2, divisao)
	} else {
		fmt.Println("Divisão por zero não é permitida!")
	}
}

