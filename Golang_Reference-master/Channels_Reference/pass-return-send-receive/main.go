package main

func main() {
	incChan := incrementor()
	cSum := puller(incChan)
	for n := range cSum {
		println(n)
	}
}

func incrementor() <-chan int { /* type channel that can only receive */
	outChan := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			outChan <- i
		}
		close(outChan)
	}()
	return outChan
}

func puller(incChan <-chan int) <-chan int { /*  Receive only channel that receives a recive only channel as well via. () */
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
