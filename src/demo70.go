package main

import (
	"fmt"
	"reflect"
)

type notKnownType struct {
	s1, s2, s3 string
}

func (n notKnownType) string() string {
	return n.s1 + "-" + n.s2 + "-" + n.s3
}

var secret interface{} = notKnownType{"Ada", "Go", "Oberon"}

func main() {
	// 返回被检查对象的实际值
	value := reflect.ValueOf(secret)
	fmt.Println(value)
	// 返回被检查对象的实际类型 main.notKnownType
	typ := reflect.TypeOf(secret)

	fmt.Println(typ)
	// 返回类型
	knd := value.Kind()
	fmt.Println(knd)
	for i := 0; i < value.NumField(); i++ {
		fmt.Println(i, value.Field(i))
	}
}
