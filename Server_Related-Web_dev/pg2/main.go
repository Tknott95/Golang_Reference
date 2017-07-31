// go run maing.go > index.html    write to html
package main

import "fmt"

func main() {
    name := "Trevor Knott"

    tpl := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
    <meta charset="UTF-8">
    <title>Hello Colorado</title>
    </head>
    <body>
    <h1>` + name + `</h1>
    </body>
    </html>`

    fmt.Println(tpl)
}
