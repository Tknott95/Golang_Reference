package main

import "fmt"

/*
  @AUTHOR - tk@trevorknott.io
  Bubble sort plucks out first two numbers compares then pushes lowest to left of array. Continues, ya know
  Here we are grabbing our first two numbers and then ++ accordingly in the loop
*/

func main() {
	var numbers []int = []int{5, 4, 7, -32, -3, 6, 4, 1, 0, 34, 234, 324423, 3, 55, 423234, 243423, 4, 455, 565}

	fmt.Println("Our list of numbers b4 sort ", numbers)
	bubbleSort(numbers)
	fmt.Println("Our list of numbers after sort ", numbers)

}

func bubbleSort(numbers []int) {
	var N = len(numbers)

	for i := 0; i < N; i++ {
		sweep(numbers)
	}
}

func sweep(numbers []int) {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1

	for secondIndex < N {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if firstNumber > secondNumber {
			// Switch the nums
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		println("Comparing the following:", firstNumber, secondNumber)

		firstIndex++
		secondIndex++
	}
}
