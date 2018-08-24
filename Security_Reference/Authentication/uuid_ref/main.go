package main

import (
	"time"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	now := time.Now().Unix() // type int64
	hourFromNow := time.Now().Add(time.Hour * 1).Unix()

	fmt.Printf("\n EXAMPLE | time.Now().Unix():\n %v \n",  now)
	fmt.Printf("\n EXAMPLE |  time.Now().Add(time.Hour + 1).Unix():\n %v \n",  hourFromNow)


	fmt.Printf("\n EXAMPLE |  bcrypt.GenerateFromPassword([]byte(input), 10) :\n %v \n",  hashInput("password"))


	
}

func hashInput(input string) []byte {
	newHash, _ := bcrypt.GenerateFromPassword([]byte(input), 10) // returns type []byte
	return newHash
}
