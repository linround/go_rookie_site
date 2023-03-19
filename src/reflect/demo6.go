package main

import (
	"fmt"
	"reflect"
)

func main() {
	//f1()
	//f2()
	//f3()
	f4()
}

func f4() {
	x := 2
	// 获取变量X对应的可取地址的值value
	d := reflect.ValueOf(&x).Elem() // d refers to the variable x

	d.Set(reflect.ValueOf(4))
	fmt.Println(x) // "4"
	// x = 3
	fmt.Println(x) // "3"

}

func f3() {
	x := 2
	// 获取变量X对应的可取地址的值value
	d := reflect.ValueOf(&x).Elem() // d refers to the variable x

	// Addr() 返回一个value ，l里面保存了指向变量的指针
	// 在value上调用interface方法，也就是返回一个interface{}
	// 使用类型断言 将得到的interface{}类型的接口 强制转为普通的类型指针

	px := d.Addr().Interface().(*float64) // px := &x
	*px = 3                               // x = 3
	fmt.Println(x)                        // "3"

}

func f2() {
	x := 2                         // value   type    variable?
	a := reflect.ValueOf(2)        // 2       int     no
	b := reflect.ValueOf(x).Elem() // 2       int     no
	c := reflect.ValueOf(&x).Elem()
	// &x      *int    no
	d := c.Elem() // 2       int     yes (x)

	fmt.Println("x", x)
	fmt.Println("a", a.CanAddr())
	fmt.Println("b", b.CanAddr())
	fmt.Println("c", c.CanAddr())
	fmt.Println("d", d.CanAddr())
}
func f1() {
	x := 2                  // value   type    variable?
	a := reflect.ValueOf(2) // 2       int     no
	b := reflect.ValueOf(x) // 2       int     no
	c := reflect.ValueOf(&x)
	// &x      *int    no
	d := c.Elem() // 2       int     yes (x)

	fmt.Println("x", x)
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
}
