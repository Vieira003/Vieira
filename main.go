package main

import (
    "fmt"
    
)

func main() {
    age := 16
    fmt.Println(age <= 20)
    fmt.Println(age >= 20)
    fmt.Println(age == 20)
    fmt.Println(age != 20)
    if age < 20 {
        fmt.Println("menor que 20 anos")
    } else if age < 10 {
        fmt.Println("menor que 10 anos")
    } else {
        fmt.Println("nao é menor que 20 anos")
    }


	names := []string{"Isadora", "Yasmim", "Martin", "Miguel", "Murilo"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continue após a posição", index, "Valor", value)
			continue
		}
		if index > 2 {
			fmt.Println("sair após", index)
			break
		}
		fmt.Println("Valor:", value)
	}
}
