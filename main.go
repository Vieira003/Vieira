package main

import (
    "fmt"
    
)

func main() {
    var saldo float32
    var saldo2 float32
    var banco int
    var soma float32
    var saldo3 float32
    var subtraçao float32
    fmt.Println("insira o valor inicial")
    fmt.Scan(&saldo)
    fmt.Println("voce deseja sacar(1) ou depositar(2) ou encerrar a sessao(3)?")
    fmt.Scan(&banco)
    if banco == 2 {
        fmt.Println("voce deseja depositar que quantia?")
        fmt.Scan(&saldo2)
        soma = saldo + saldo2
        fmt.Println("o seu saldo total é", soma)

    }
    if banco == 1 {
        fmt.Println("voce deseja sacar que quantia?")
        fmt.Scan(&saldo3)
        subtraçao = saldo - saldo3
        fmt.Println("o seu saldo total é", subtraçao)
    }
    if banco == 3 {
        fmt.Println("sessao encerrada")
    }
    



}