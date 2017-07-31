package main

import (
    "os"
    "text/template"
    "log"
)

var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseFiles("tpl.gothml"))
}

func main() {
    xs := []string{"zero", "one", "two", "three", "four", "five"}

    data := struct {
        Words []string
        Lname string
    }{
        xs,
        "McLeod",
    }

    err := tpl.Execute(os.Stdout, data)
    if err != nil {
        log.Fatalln(err)
    }
}