package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 12.0
	fmt.Println(number)
	fmt.Println(reflect.TypeOf(number))
}
