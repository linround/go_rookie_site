package main

import "fmt"

func main() {
	var x int
	fmt.Println(&x == nil)
	fmt.Println(&x)
}
