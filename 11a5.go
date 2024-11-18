package main

import "fmt"

func main() {
	var name string
	var surname string
	var age int
	fmt.Println("Start...")
	fmt.Scanln(&name, &surname, &age)
	fmt.Println("Your name is", name, surname, "and you are", age, "years old")
	fmt.Println("End...")
}
