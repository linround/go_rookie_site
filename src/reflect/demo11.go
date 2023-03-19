package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a, b []string = nil, []string{}
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println(reflect.DeepEqual(a, b)) // "false"

	var c, d map[string]int = nil, make(map[string]int)
	fmt.Println("c:", c)
	fmt.Println("d:", d)
	fmt.Println(reflect.DeepEqual(c, d)) // "false"

}
