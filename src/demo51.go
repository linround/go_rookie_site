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

	//	map类型的切片
	//	第一次分配切片，第二次分配切片中每个map元素

	items := make([]map[int]int, 5)
	for i := range items {
		items[i] = make(map[int]int, 1)
		items[i][2] = 100
		items[i][0] = 10
	}
	fmt.Println(items)
}
