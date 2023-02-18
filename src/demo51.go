package main

import "fmt"

func main() {
	var mapList map[string]int
	var mapAssigned map[string]int
	mapList = map[string]int{
		"one": 1,
		"two": 2,
	}
	mapAssigned = mapList
	mapAssigned["key"] = 1000

	// 以下可以证明，map是引用传递
	fmt.Println(mapList)
	fmt.Println(mapAssigned)

}
