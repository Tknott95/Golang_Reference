package main

import (
    "io"
    "net/http"
)

func dPage(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, "Another page about animals.. and yes DOG PAGE")
}

func cPage(res http.ResponseWriter, req *http.Request) {
    io.WriteString(res, "I am allergic and hate cats! Cats suck!!!!!!!")
}

func main() {

    http.HandleFunc("/dog/", dPage)
    http.HandleFunc("/cat/", cPage)

    http.ListenAndServe(":8080", nil)
}
