package main

import (
    "io"
    "net/http"
)

type tkDogpage int

func (dt tkDogpage) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, "Dog page bruh bruh")
}

type tkCatPage int

func (ct tkCatPage) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, "Cat page EWWW :(")
}

func main() {
    var dt tkDogpage
    var ct tkCatPage

    mux := http.NewServeMux()
    mux.Handle("/dog/", dt)
    mux.Handle("/cat/", ct)

    http.ListenAndServe(":8080", mux)
}
