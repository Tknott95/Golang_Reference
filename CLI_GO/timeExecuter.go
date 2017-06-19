package main

import (
  "fmt"
  "os"
  "time"
)

func main() {
  now := time.Now()
  day := now.Weekday()

  hr, min, sec := now.Clock()

  // Changes testing spec based on day

  fmt.Println("Code will only give golden ticket at right time!")

 if day == 4 && hr == 10 && min == 40 && sec == 0 {
  fmt.Println("GOLDEN TICKET, YOU GET $$$$")
 } else {
  os.Exit(0)
 }
}
