package main

import (
"fmt"
)



func main() {
    alunoIdade := make(map[string]int)
    alunoIdade["bruno"] = 16
    alunoIdade["murilo"] = 16
    alunoIdade["fabiano"] = 40
    alunoIdade["manu"] = 15
    fmt.Println("idade do Bruno",alunoIdade["bruno"])
    notasAlunos := map[string]float64{
        "bruno" : 9.7,
        "otavio" : 10, 
        "fabiano" : 8.7,
        "manu" : 9.5,  
 }
 for nome,nota := range notasAlunos{
    fmt.Printf("%s tirou a nota %.1f \n", nome, nota)
 }
}