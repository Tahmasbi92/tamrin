package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 12.25
	fmt.Println(number)
	fmt.Println(reflect.TypeOf(number))
}
