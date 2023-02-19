package main

import (
	"fmt"
	"reflect"
)

type myInit int

func main() {
	var x float64 = 3.4
	fmt.Println("type", reflect.TypeOf(x))
	//传递一个x的拷贝创建了v
	v := reflect.ValueOf(x)
	fmt.Println("CanSet:", v.CanSet())
	fmt.Println("value:", v)
	fmt.Println("Type:", v.Type())
	fmt.Println("Kind:", v.Kind())
	fmt.Println("Float:", v.Float())
	fmt.Println("Interface:", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)
	println("main69**********")
	main69()

}

func main69() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("CanSet:", v.CanSet())
	v = reflect.ValueOf(&x)
	fmt.Println("CanSet:", v.CanSet())
	fmt.Println("type:", v.Type())
	// 使用Elem 来使得反射可以被设置
	v = v.Elem()
	fmt.Println(v)
	fmt.Println("CanSet:", v.CanSet())
	v.SetFloat(3.1415)
	fmt.Println(v.Interface())
	// 最终 v值和x值都发生变化
	fmt.Println(v)
	fmt.Println(x)

}
