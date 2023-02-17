package main

import "fmt"

var avv string

func console(v int) {
	fmt.Println(v)
}

func main() {
	avv = "G"
	print(avv)
	f1()
	var name = 5 / 4
	fmt.Println(name)
	var s int = 100
	s++
	console(s)
}
func f1() {
	avv := "o"
	print(avv)
	f2()
}
func f2() {
	print(avv)
}
