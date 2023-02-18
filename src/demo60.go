package main

import "fmt"

type TwoInt struct {
	a, b int
}

func (tn *TwoInt) add() int {
	return tn.b + tn.a
}
func (tn *TwoInt) addParam(p int) int {
	return tn.a + tn.b + p
}

func main() {
	two := TwoInt{1, 5}
	fmt.Println(two.add())
	fmt.Println(two.addParam(4))
}
