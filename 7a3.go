package main

import (
	"fmt"
	"reflect"
)

func main() {
	full_name := "Reza Habibi"
	fmt.Println(full_name)
	fmt.Println(reflect.TypeOf(full_name))
}
