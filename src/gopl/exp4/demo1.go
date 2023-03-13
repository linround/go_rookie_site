package main

import "fmt"

func main() {
	ages := map[string]int{"name": 1}
	delete(ages, "name")
	fmt.Println(ages["name"])
	// len 为4  cap为8
	names := make([]int, 4, 8)
	names[3] = 5
	n2 := names[2]
	fmt.Println(n2)
	fmt.Println(len(names), cap(names))
	fmt.Println(names)
}
