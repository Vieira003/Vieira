package main 

import "fmt" 

func main (){
  var ages = [4]int{11,8,16,14}
  nomes :=[4]string{"Rosa", "Amy", "Jack", "miguel"}
  fmt.Println(ages)
  fmt.Println(nomes)
//   slice
var score = []int{100, 200, 300,}
fmt.Println(score)
rangeOne := score[1:3]
fmt.Println(rangeOne)
rangeTwo := score[2:]
fmt.Println(rangeTwo)
rangeTree := score[:3]
fmt.Println(rangeTree)

var superherois = []string{"Deadpool", "Homem Aranha", "Hulk", "Thor"}
fmt.Println(superherois)
superherois = append(superherois, "Homem de Ferro")
fmt.Println(superherois, len(superherois), cap(superherois))


}



