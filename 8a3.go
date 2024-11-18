package main

import (
	"fmt"
	"reflect"
)

func main() {
	var is_male bool = true
	fmt.Println(is_male)
	fmt.Println(reflect.TypeOf(is_male))
}
