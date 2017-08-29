package main

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := -10; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for f := 44; f > 33; f-- {
			c <- f
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for val := range c {
		println(val)
	}
}
