package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"math"
	"math/rand"
)

var flagEx int

func init() {

	flag.IntVar(&flagEx, "f", 1234, "help msg for flag example")
	flag.Parse() /* In order to allow flag to overwrite default */
}

func main() {
	r := rand.New(rand.NewSource(77))
	h := sha256.New()

	h.Write([]byte("hello world\n"))
	fmt.Printf("%X", h.Sum(nil))   /* "%x or %X controls base 16, lower-case, two characters per byte" */
	fmt.Printf("\n%x", h.Sum(nil)) /* "%x or %X controls base 16, lower-case, two characters per byte" */

	fmt.Println("\nYour example flag is:", flagEx)
	fmt.Println("Random floating points (64) (32):", r.Float64(), r.Float32())

	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))

}
