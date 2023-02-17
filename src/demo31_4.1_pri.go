package main

var av string = "G"

func main() {
	n()
	m()
	n()
}
func n() {
	print(av)
}
func m() {
	av = "o"
	print(av)
}
