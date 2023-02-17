package main

import "fmt"

func main() {
	kk("1", "5", "8", "9")
}

func kk(s ...string) {
	fmt.Print(s, "kk")
	kkk(s...)
	kkk2(s)
}
func kkk(s ...string) {
	fmt.Print(s, "kkk")
}
func kkk2(s []string) {
	fmt.Print(s)
}
