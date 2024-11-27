package main

import (
	"fmt"
)

func main() {
	var nameee = "123"
	// 012
	for number, sina := range nameee {
		fmt.Println(number)
		fmt.Println(sina)
	}
}
