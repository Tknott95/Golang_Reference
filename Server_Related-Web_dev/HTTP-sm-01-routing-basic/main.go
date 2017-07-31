package main

import (
    "io"
    "net/http"
)

type tkServerBruh int

func (m tkServerBruh) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    switch req.URL.Path {
        case "/ollie":
            io.WriteString(w, "Ollie the dog reigns King!")
        case "/ugly-cat":
            io.WriteString(w, "Who cares?... To many cats on the web :(")
    }
}

func main() {
    var goBro tkServerBruh
    http.ListenAndServe(":8080", goBro)
}
