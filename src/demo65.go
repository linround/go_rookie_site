package main

import "fmt"

type s struct {
	side float32
}
type c struct {
	r float32
}

type sp interface {
	area() float32
}

func main() {
	var areaInt sp
	sq1 := new(s)
	sq1.side = 5
	areaInt = sq1
	t, ok := areaInt.(*s)
	fmt.Println(t, ok)

}

func (t *s) area() float32 {
	return t.side
}
func (d *c) area() float32 {
	return d.r
}
