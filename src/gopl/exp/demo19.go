package main

import "fmt"

func remove(slice []int, i int) []int {
	fmt.Println(slice)
	// copy 函数 将后面的子slice
	fmt.Println(slice[i:])
	fmt.Println(slice[i+1:])
	copy(slice[i:], slice[i+1:])
	fmt.Println(slice)
	return slice[:len(slice)-1]
}
func main() {
	s := []int{5, 6, 7, 8, 9}
	r := remove(s, 1)
	fmt.Println(s)
	fmt.Println(r)
}
