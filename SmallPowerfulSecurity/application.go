package main

import (
	"fmt"
	"net/http"

	analytics "github.com/tknott95/Golang_Reference/SmallPowerfulSecurity/Controllers/analyticsCtrl"
)

func landing(w http.ResponseWriter, req *http.Request) {
	fmt.Println(`Dam nice middleware ;)`)
	w.Write([]byte("Hello"))

}

func panicSimulator(w http.ResponseWriter, req *http.Request) {
	panic(analytics.ErrInvalidEmail)
}

func main() {
	logger := analytics.CreateLogger("secretapi")
	fmt.Println("Initialized... Do you even fibbonacci bro? 👽")

	http.Handle("/", analytics.Time(logger, landing))
	http.Handle("/panic", analytics.Recover(panicSimulator))
	http.ListenAndServe(":8080", nil)
}
