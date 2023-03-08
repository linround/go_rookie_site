package main

import (
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "df", "separator")

func main() {
	flag.Parse()
	fmt.Println("n:", *n)
	fmt.Println("sep:", *sep)
	fmt.Println(flag.Args())
	fmt.Println("n:", *n)
	if *n {
		fmt.Println("n满足")
	}

	fmt.Println(5 / 8) // 取整
	fmt.Println(5 % 8) // 取余

}
