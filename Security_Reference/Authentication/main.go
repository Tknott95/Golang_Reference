package main

import (
	"fmt"
	"time"

	"github.com/satori/go.uuid"
	ctrlJwt "github.com/tknott95/Golang_Reference/Security_Reference/Authentication/JWT"
	"golang.org/x/crypto/bcrypt"
)

// bcrypt.CompareHashAndPassword([]byte(testUser.Password), []byte(user.Password)) == nil
// }

func main() {
	timeExamples()

	ctrlJwt.InitKeys()

	printing()

}

func printing() {
	z, _ := ctrlJwt.GenerateJWT(newU4(), "tk")
	fmt.Printf("\n EXAMPLE |  bcrypt.GenerateFromPassword([]byte(input), 10) :\n %x \n", hashInput("password"))

	fmt.Printf("\n EXAMPLE |  uuid.Must(uuid.NewV4()) :\n %s \n", newU4())

	fmt.Printf("\n EXAMPLE | ctrlJwt.GenerateJWT(newU4(), tk) : \n  %s \n", z)
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
