package main

import (
	"fmt"
	"log"
	"runtime"
)

func f() (r int) {
	defer func() {
		r++
	}()

	return 1
}
func main() {
	fmt.Println(f())
	where()
}

var where = func() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s,%d", file, line)
}
