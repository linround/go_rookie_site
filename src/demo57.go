package main

import (
	"fmt"
	"reflect"
)

type Foo map[string]string
type Bar struct {
	one string "on 1 2 3"
	two int
}

func main() {
	y := new(Bar)
	y.two = 15
	(*y).one = "kl"
	yyType := reflect.TypeOf(*y)
	yyFiled := yyType.Field(0)
	fmt.Println(yyFiled)

	//	可以使用make的三种类型 slice/map/channel
	x := make(Foo)
	x["1000"] = "1000"

	fmt.Println(y, x)
}
