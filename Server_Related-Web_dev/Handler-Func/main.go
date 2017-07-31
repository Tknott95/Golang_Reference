package main

import (
	"io"
	"net/http"
)

func ds(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "DOG PAGE")
}

func cs(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "CAT PAGE")
}

func main() {

	http.Handle("/dog", http.HandlerFunc(ds))
	http.Handle("/cat", http.HandlerFunc(cs))

	http.ListenAndServe(":8080", nil)
}
