package main

import "fmt"

func main() {
	i := 1
	fmt.Println("Start Loop...")
	for i <= 10 {

		if i == 6 {
			break
		}
		fmt.Println(i)
		i += 1
	}
	fmt.Println("End Loop...")
}
