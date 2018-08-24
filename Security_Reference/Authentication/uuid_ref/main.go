package main

import (
	"time"
	"fmt"
)

func main() {

	now := time.Now().Unix() // type int64

	fmt.Printf("\n EXAMPLE | time.Now().Unix():\n %v \n",  now)
}