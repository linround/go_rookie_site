package main

import "strconv"

func main() {
	println(strconv.IntSize)

	var orig string = "666"
	var an int
	var newS string
	an, _ = strconv.Atoi(orig)
	println(an)
	an += 5
	newS = strconv.Itoa(an)
	println(newS)
}
