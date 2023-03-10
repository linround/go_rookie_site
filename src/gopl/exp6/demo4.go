package main

import "fmt"

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}
func main() {
	var b ByteCounter
	b.Write([]byte("hello"))
	fmt.Println(b)
}
