package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(1))
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}

}
