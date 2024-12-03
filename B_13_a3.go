package main

import "fmt"

func main() {
	numbers := []int{2, 3, 5, 11, 13, 1}
	// LENGTH        1  2  3   4   5  6
	// len مخفف length یعنی طول می باشد که در این شکل 6 تا می باشد
	fmt.Println(len(numbers))
}
