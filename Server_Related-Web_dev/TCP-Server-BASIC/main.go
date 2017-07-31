package main

import (
    "fmt"
    "io"
    "log"
    "net"
)

func main() {
    li, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalln(err)
    }
    defer li.Close()

    for {
        conn, err := li.Accept()
        if err != nil {
            log.Println(err)
            continue
        }

        io.WriteString(conn, "\nTCP SERVER connectione established...\n")
        fmt.Fprintln(conn, "Hello, I am a TCP server!")
        fmt.Fprintln(conn, "%v", "Hope all is well bruh bruh!")

        conn.Close()
    }
}
