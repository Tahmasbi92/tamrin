package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "30"
	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))
}
