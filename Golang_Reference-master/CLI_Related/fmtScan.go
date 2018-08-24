package main

import (
  "fmt"
)

func main() {
 var name string
 var zipCode string

 fmt.Println("What is your name?")
 fmt.Scanf("%s", &name)
 
 fmt.Println("Welcome,", name, "to aquire weather enter your Zip Code below")
 fmt.Scanf("%s", &zipCode) 

 fmt.Println("Location:", zipCode)
}
