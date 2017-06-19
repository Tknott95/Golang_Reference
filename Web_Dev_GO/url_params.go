package main

import (
  "io"
  "net/http"
)

func main() {
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter , req *http.Request) {
 ex := req.FormValue("example")
 ex2 := req.FormValue("e")

 io.WriteString(w, "?example=suh%20bruh&e=example2")
 io.WriteString(w, "\nExample param is: "+ex)
 io.WriteString(w, "\nExample param is: "+ex2)
}
