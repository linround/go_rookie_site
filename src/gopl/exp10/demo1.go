package main

import (
	"fmt"
	"rookie/src/gopl/pkg2"
	"rookie/src/gopl/pkgDemo"
)

// pkgDemo init 执行
// pkg2 init
// pkgDemo.Test()
func main() {
	pkgDemo.Test()
	pkg2.Test()
	fmt.Println("main vvvvvvv", pkgDemo.V)
}
