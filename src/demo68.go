package main

type absInterface interface {
	abs()
}
type squareInterface interface {
	sqr()
}
type point struct {
}

func (p point) abs() {

}
func (p point) sqr() {

}

func main() {
	var ai absInterface
	var si squareInterface
	pp := new(point)
	var empty interface{}
	empty = pp
	ai = empty.(absInterface)
	si = ai.(squareInterface)
	empty = si
	ai.abs()
}
