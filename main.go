package main

import (
	"fmt" 
	"strings"
	"sort"
)

func main (){
  greenting := "Hello my friends!"
  fmt.Println(strings.Contains(greenting,"friends"))
fmt.Println(strings.ReplaceAll(greenting, "Hello" , "Hi"))
fmt.Println(strings.ToUpper(greenting))
fmt.Println(strings.Index(greenting, "my"))
fmt.Println(strings.Split(greenting, "friends"))
ages := []int {50, 80 ,60}
sort.Ints(ages)
fmt.Println(ages)
index := sort.SearchInts(ages, 50)
fmt.Println(index)
names := []string{"Caruzo", "Cris", "Julius"}
sort.Strings(names)
fmt.Println(names)
fmt.Println(sort.SearchStrings(names, "Cris"))





}



