package main

import "fmt"

func main() {
	numbers := []int{2, 3, 5, 11, 13, 1}
	//   0  1  2   3   4  5

	// Change a one Index in Slices
	fmt.Println(numbers)
	numbers[0] = 6266
	fmt.Println(numbers)
}
