package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Reza"
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
}
