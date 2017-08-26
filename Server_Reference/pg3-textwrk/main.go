package main

import (
    "log"
    "os"
    "text/template"
)

func main() {
    tpl, err := template.ParseFiles("tpl.gohtml") // i parse files give it file name   - takes 0 or more strings in params.. going to parse then store with pointer to template (Container holding all parsed templates by me)
    if err != nil {
        log.Fatalln(err)
    }
    nf, err := os.Create("indexOScreated.html")
    if err != nil {
        log.Println("error creating file", err)
    }
    defer nf.Close()
    
    err = tpl.Execute(nf, nil) // execute takes writer and data
    if err != nil {
        log.Fatalln(err)
    }
}
