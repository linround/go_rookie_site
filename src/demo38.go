package main

import (
	"os"
	"strconv"
)

func main() {
	var orig string = "789"

	an, err := strconv.Atoi(orig)
	if err != nil {
		println("出错了", err)
	}
	println(an)
	main38()
}

func main38() {
	f, err := os.Open("./demo37")
	if err != nil {
		println(err, "出错了")
	}
	println(f)
}
