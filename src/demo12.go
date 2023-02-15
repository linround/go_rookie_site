package main

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	next := getSequence()
	println(next())
	println()
	println(next())
	println()
	println(next())
	println()
	println(next())
	println()
}
