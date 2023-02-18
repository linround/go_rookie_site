package main

import (
	"fmt"
	"sort"
)

func main() {
	slForm := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slForm)
	fmt.Println(slTo)
	fmt.Println(n)
	sl3 := []int{1, 2, 3, 50, 4, 5}
	sl4 := append(sl3, 100, 200)
	fmt.Println(sl3)
	fmt.Println(sl4)
	s := "hello"
	c := []byte(s)
	fmt.Println(c)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)

	b := "s" > "a"
	fmt.Println(b)
	sort.Ints(sl3)
	fmt.Println(sl3)

}
