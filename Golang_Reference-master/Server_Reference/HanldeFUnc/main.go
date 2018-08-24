package main

import (
	"io"
	"net/http"
)

func dww(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "DOGS 4 LIFEE!")
}

func caa(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "I HATE CATS")
}

func main() {

	http.HandleFunc("/dog", dww)
	http.HandleFunc("/cat", caa)

	http.ListenAndServe(":8080", nil)
}
