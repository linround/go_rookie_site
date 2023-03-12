package main

import (
	"fmt"
	"math"
)

func main() {
	var f float32 = 1 << 10
	fmt.Println(f)
	fmt.Println(f+1 == f)
	fmt.Println(f+2 == f)
	fmt.Printf("%g \n", f)
	fmt.Printf("%e \n", f)
	fmt.Printf("%f \n", f)
	m := math.NaN()
	fmt.Println(math.IsNaN(m))
}
