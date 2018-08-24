package main

import (
    "bufio"
    "fmt"
    "strings"
)

func main() {
    s := "Lorem ipsum dolor sit amet\n, consectetur adipiscing elit, sed do ei\nusmod tempor incididu\nnt ut labore et dolore mag\nna aliqua. Ut enim ad minim\n veniam, quis nostrud exercitation\n ullamco laboris nisi ut aliquip \nex ea commodo consequat. Duis aute ir\nure dolor in reprehend\nerit in voluptat\ne velit esse"

    scanner := bufio.NewScanner(strings.NewReader(s))

    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }

}