package main

import (
	"sync"
)

/* tk@trevorknott.io */
func main() {
	c := make(chan int)

	var _wg sync.WaitGroup
	_wg.Add(2)
	/*   SHOWED A RACE CONDITION IF THE WAIT WAS IN THE funcs */
	/*  This was due to concurrent programming and setting a value at the same time to this one va aka _wg fixes made now */
	go func() {
		/*  Anonymous Function */
		for i := 0; i < 10; i++ {
			c <- i
		}
		_wg.Done()
	}()

	go func() {
		/*  Anonymous Function */
		for i := 10; i < 20; i++ {
			c <- i
		}
		_wg.Done()
	}()

	go func() {
		/* This is also a channel running and just waiting fr the other channels to be finished (looping) */
		_wg.Wait()
		close(c)
	}()

	for channelVal := range c {
		println(channelVal)
	}
}
