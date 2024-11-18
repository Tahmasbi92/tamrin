package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := -123
	fmt.Println(number)
	fmt.Println(reflect.TypeOf(number))
}
