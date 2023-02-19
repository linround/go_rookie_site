package main

import "fmt"

type iDuck interface {
	quack()
	walk()
}

func duckDance(d iDuck) {
	d.quack()
	d.walk()
}

type bird struct {
}

func (b *bird) quack() {
	fmt.Println("quacking")
}
func (b *bird) walk() {
	fmt.Println("walking")
}

func main() {
	b := new(bird)
	duckDance(b)
}
