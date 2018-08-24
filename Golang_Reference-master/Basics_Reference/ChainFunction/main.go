package main

import "fmt"


func main() {
	fmt.Println("Initialized... Do you even fibbonacci bro? 👽")
	variaticFunction()
	chainFunction()
}

func variaticFunction() {
	sum := Add(1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89)
	fmt.Println(sum)
}

func chainFunction() {
	chain := &Chain{0}
	sum2 := chain.AddNext(1).AddNext(1).AddNext(2).AddNext(3).AddNext(5).AddNext(8).AddNext(13).AddNext(21).AddNext(34).AddNext(55).AddNext(89).Finally(0)
	fmt.Println(sum2)
}
