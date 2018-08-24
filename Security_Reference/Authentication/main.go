package main

import (
	"fmt"
	"time"

	"github.com/satori/go.uuid"
	ctrlJwt "github.com/tknott95/Golang_Reference/Security_Reference/Authentication/JWT"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	timeExamples()

	ctrlJwt.Init()

	fmt.Printf("\n EXAMPLE |  bcrypt.GenerateFromPassword([]byte(input), 10) :\n %x \n", hashInput("password"))

	fmt.Printf("\n EXAMPLE |  uuid.Must(uuid.NewV4()) :\n %s \n", newU4())

}

func timeExamples() {
	now := time.Now().Unix()                               // type int64
	minsFromNow := time.Now().Add(time.Minute * 20).Unix() // type int64
	hourFromNow := time.Now().Add(time.Hour * 1).Unix()    // type int64

	print("TIME EXAMPLES")
	fmt.Printf("\n EXAMPLE | time.Now().Unix():\n %v \n", now)
	fmt.Printf("\n EXAMPLE |  time.Now().Add(time.Minute * 20).Unix():\n %v \n", minsFromNow)
	fmt.Printf("\n EXAMPLE |  time.Now().Add(time.Hour + 1).Unix():\n %v \n", hourFromNow)
	print("ENDING TIME EXAMPLES \n")
	return
}

func hashInput(input string) []byte {
	newHash, _ := bcrypt.GenerateFromPassword([]byte(input), 10) // type []byte
	return newHash
}

func newU4() uuid.UUID {
	u4 := uuid.Must(uuid.NewV4()) // type uuid.UUID  ( under hood is [size]byte )
	return u4
}
