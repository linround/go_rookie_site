package main

import (
	"fmt"

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

func main() {
	two := TwoInt{1, 5}
	fmt.Println(two.add())
	fmt.Println(two.addParam(4))
	fmt.Println(two.a)

	p := new(person.Person)
	p.FirstName()
}
