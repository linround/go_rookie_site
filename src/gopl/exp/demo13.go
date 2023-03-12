package main

import "fmt"

func main() {

	s := "abc"
	b := []byte(s)
	b[1] = 2
	s2 := string(b)
	fmt.Println(s2)
	fmt.Println(s)
}
