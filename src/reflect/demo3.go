package main

import (
	"gopl.io/ch12/display"
	"rookie/src/test/eval"
)

func main() {
	f1()
}
func f1() {
	e, _ := eval.Parse("sqrt(A/pi)")
	display.Display("e", e)
}
