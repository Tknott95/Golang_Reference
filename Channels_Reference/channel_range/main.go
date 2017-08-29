package main

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i /*  Channel receiving i */
		}
		close(c)
	}()

	for valueFromChannel := range c {
		println(valueFromChannel)
	}
}
