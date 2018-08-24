package main

import (
    "io"
    "net/http"
)

func main() {

    http.HandleFunc("/", pictureExample)
    http.ListenAndServe(":8080", nil)

}

 func pictureExample(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")

        io.WriteString(w, `
                <img src="http://debatepost.com/wp-content/uploads/2016/07/anonymous-1200.png">
                `)
    }