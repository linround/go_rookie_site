package main

import "fmt"

type Point struct{ X, Y int }

func main() {

	p := Point{1, 2}
	q := Point{1, 1}
	fmt.Println(p.X == q.X) // "true"
	fmt.Println(p == q)     // "false"

}
