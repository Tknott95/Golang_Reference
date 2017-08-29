package main

func main() {
	incChan := incrementor()
	cSum := puller(incChan)
	for n := range cSum {
		println(n)
	}
}

func incrementor() chan int {
	outChan := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			outChan <- i
		}
		close(outChan)
	}()
	return outChan
}

func puller(incChan chan int) chan int {
	out := make(chan int)
	go func() {
		var val int
		for v := range incChan {
			val += v
		}
		out <- val
		close(out)
	}()
	return out
}
