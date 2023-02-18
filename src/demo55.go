package main

import (
	"./uc"
	"fmt"
)

type St struct {
	name string
	_    int
}

type MySt struct {
	i int
}

var ms MySt

var st St

func Text() {
	fmt.Println("demo55")
}

type Inter struct {
	start int
	end   int
}

func main() {
	intr := Inter{
		100, 200,
	}
	fmt.Println(intr)
	s := "ac"
	st.name = "name"
	ms = MySt{10}
	fmt.Println(ms.i)
	r := uc.UpperCase(s)
	fmt.Println(r)
}
