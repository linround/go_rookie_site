package main

import "fmt"

func main() {
	callback(10, add)
}
func callback(a int, f func(int2 int, int3 int) int) {
	v := f(1, 5)    // 6
	fmt.Println(v)  // 6
	v2 := a + v     // 16
	fmt.Println(v2) // 16
}
func add(int2 int, int3 int) (v int) {
	v = int2 + int3
	return
}
