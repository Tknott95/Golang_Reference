package main

/*  Chain */
type Chain struct {
	Sum int
}

/*  Can bring in as many ints as I want, Variatic Function */
func Add(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func (c *Chain) AddNext(num int) *Chain {
	c.Sum += num
	return c
}

func (c *Chain) Finally(num int) int {
	return c.Sum + num
}
