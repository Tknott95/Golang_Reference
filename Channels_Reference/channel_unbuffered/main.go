package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i /*  on to the channel */
		}
	}()
	/* These two channels are communicating w/ goroutines */
	go func() {
		for {
			fmt.Println(<-c) /*  from channel give me value */
		}
	}()
	time.Sleep(time.Second)
}
