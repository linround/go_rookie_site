package main

import "runtime"

func main() {
	var intP *int
	var i1 int
	intP = &i1
	println(intP)
	println(i1)
	if runtime.GOOS == "windows" {
		println("sh")
	}

	var cond int
	cond = 100
	if cond > 10 {
		println("cond", cond)
	}
}
