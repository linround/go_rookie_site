package main

import (
	"fmt"
	"strconv"

	"./person"
)

type TwoInt struct {
	a, b int
}

// add 接收一个指向TwoInt的指针

// 去掉*号就是 add通过拷贝接受TwoInt的值

func (tn *TwoInt) add() int {
	tn.a = 1000
	return tn.b + tn.a
}

func (tn *TwoInt) addParam(p int) int {
	return tn.a + tn.b + p
}
func (tn *TwoInt) String() string {
	return "(" + strconv.Itoa(tn.a) + "/888" + strconv.Itoa(tn.b) + ")"
}

func main() {
	two := TwoInt{1, 5}
	// 这里需要传递引用类型的
	fmt.Printf("%v", &two)

	// 这里的.号是自动解引特性 不需要传递地址
	// 会自动解析地址的
	fmt.Println(two.add())
	fmt.Println(two.addParam(4))
	fmt.Println(two.a)

	p := new(person.Person)
	p.FirstName()
}
