package main

import (
	"fmt"
	"time"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	now := time.Now().Unix() // type int64
	hourFromNow := time.Now().Add(time.Hour * 1).Unix()

	fmt.Printf("\n EXAMPLE | time.Now().Unix():\n %v \n", now)
	fmt.Printf("\n EXAMPLE |  time.Now().Add(time.Hour + 1).Unix():\n %v \n", hourFromNow)

	fmt.Printf("\n EXAMPLE |  bcrypt.GenerateFromPassword([]byte(input), 10) :\n %x \n", hashInput("password"))

	fmt.Printf("\n EXAMPLE |  uuid.Must(uuid.NewV4()) :\n %s \n", newU4())

}

func hashInput(input string) []byte {
	newHash, _ := bcrypt.GenerateFromPassword([]byte(input), 10) // returns type []byte
	return newHash
}

func newU4() uuid.UUID {
	u4 := uuid.Must(uuid.NewV4())
	return u4
}
