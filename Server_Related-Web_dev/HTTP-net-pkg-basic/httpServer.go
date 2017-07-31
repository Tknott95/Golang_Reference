package main

import (
    "fmt"
    "net/http"
)

type shamrock int

func (s shamrock) ServeHTTP(wrx http.ResponseWriter, req *http.Request) {
    wrx.Header().Set("Mcleod-Key", "this is from mcleod")
    wrx.Header().Set("Content-Type", "text/html; charset=utf-8")
    fmt.Fprintln(wrx, "<h1>At the end of the rainbow is not gold, it is Connor McGregor coming to snap kick yo face!</h1>")
}

func main() {
    var tk shamrock
    http.ListenAndServe(":8080", tk)
}