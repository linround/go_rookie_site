package main

func main() {
	var x []int = []int{1, 5}
	go func() { x = make([]int, 10) }()
	go func() { x = make([]int, 1000000) }()
	x[1] = 1 //

}
