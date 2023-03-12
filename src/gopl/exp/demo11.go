package main

import "fmt"

func main() {
	s := "hello world"
	fmt.Println(s[0])
	fmt.Println(string(s[0]))
	fmt.Printf("%c \n", s[0])
	fmt.Printf("%q \n", s[0])
}
