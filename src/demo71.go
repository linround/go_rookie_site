package main

import (
	"fmt"
	"reflect"
)

type tt struct {
	a int
	b string
}

func main() {
	// t是一个值
	t := tt{22, "lin"}
	// 指针反射 并开始修改
	s := reflect.ValueOf(&t).Elem()
	typeofT := s.Type()
	fmt.Println(typeofT)
	fmt.Println(s.CanSet())
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Println(i,
			typeofT.Field(i).Type, f, f.Type())
	}
	fmt.Println(t)
	//
	//s.Field(0).SetInt(1000)
	//s.Field(1).SetString("yuan")
}
