package main

import (
  "flag"
  "fmt"
)

func mult(x float64, y float64) float64 {
  return x * y
}

func main() {
 varPointer := flag.Float64("n", 52.0287, "a float")
  
 flag.Parse()

 fmt.Println("Starting Number:", *varPointer)

 fmt.Println("Final", mult(*varPointer, 3.14))
}
