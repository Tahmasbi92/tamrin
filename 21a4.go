package main

import "fmt"

func main() {
	i := 1
	fmt.Println("Start Loop...")
	for i <= 5 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println("End Loop...")
}
