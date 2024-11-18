package main

import (
	"fmt"
	"reflect"
)

func main() {
	number := 30
	fmt.Println(number)
	fmt.Println(reflect.TypeOf(number))
}
