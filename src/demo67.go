package main

import "fmt"

var i = 5
var str = "abc"

type person struct {
	name string
	age  int
}

type any interface {
}

func main() {
	var v any
	v = 5
	fmt.Println(v)
	v = str
	fmt.Println(v)
	p := &person{"lin", 15}
	p.age = 105
	v = p

	switch t := v.(type) {
	case int:
		println("int", t)
	case string:
		println("string")
	case *person:
		fmt.Println("person", t)
	default:
		fmt.Println(t)
	}
}
