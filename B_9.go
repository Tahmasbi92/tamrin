package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter something: ")
	scanner.Scan()
	userInput := scanner.Text()
	// Try to parse the user input as a number
	if num, err := strconv.Atoi(userInput); err == nil {
		fmt.Printf("%d is an integer\n", num)
	} else if _, err := strconv.ParseFloat(userInput, 64); err == nil {
		fmt.Printf("%s is a float\n", userInput)
	} else {
		fmt.Printf("%s is a string\n", userInput)
	}
}
