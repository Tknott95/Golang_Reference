package main

func main() {
	c := factorial(4)

	for val := range c {
		println("FacVal: ", val)
	}
}

func factorial(n int) chan int {

	outChan := make(chan int)
	go func() {
		total := 1
		for i := n; i > 0; i-- {
			total *= i
			outChan <- total
		}
		close(outChan)
	}()
	return outChan
}
