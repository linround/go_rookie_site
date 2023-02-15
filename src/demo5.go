package main

func main() {
	const (
		a = iota
		b
		c
		d
		e
		f = 100
		g
		h = 78
		i = iota
	)
	println(a, b, c, d, e, f, g, h, i)
}
