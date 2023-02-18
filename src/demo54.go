package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(1)
	c := big.NewInt(math.MaxInt)
	e := big.NewInt(0)
	a.Add(a, c) // c+0
	b.Add(b, c) // c+1
	e.Sub(b, a) // b-a
	fmt.Println("e", e)

	d := big.NewInt(0)
	fmt.Println(c)
	c.Add(c, a)
	d.Add(d, b)
	e.Sub(c, d)
	a.Sub(a, b)
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)

	rm := big.NewRat(math.MaxInt64, 1956)
	fmt.Println(rm)
	rn := big.NewRat(-1956, math.MaxInt64)
	fmt.Println(rn)
	ro := big.NewRat(19, 56)
	fmt.Println(ro)

}
