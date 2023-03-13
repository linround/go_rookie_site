package main

import "fmt"

func main() {
	var a [4]int = [4]int{1: 5}
	zero(&a)
	for i, v := range a {
		fmt.Println(i, ":", v)
	}
}
func zero(ptr *[4]int) {
	//for i := range ptr {
	//	ptr[i] = 0
	//}
	*ptr = [4]int{}
}
