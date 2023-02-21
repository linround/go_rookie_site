package main

import "fmt"

func main() {
	map1 := map[string]int{"one": 1, "two": 2}
	println("遍历map")
	for key, val := range map1 {
		fmt.Println(key, val)
	}
	fmt.Println("检查key值是否存在")
	_, isPresent1 := map1["three"]
	if !isPresent1 {
		fmt.Println("three 不存在")
	}
	v2, isPresent2 := map1["one"]
	if isPresent2 {
		fmt.Println("存在 one")
		fmt.Println(v2)
	}
}
