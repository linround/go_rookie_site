package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("Not found error")

func main() {
	fmt.Printf("error: %v", errNotFound)
	errs := fmt.Errorf("math: square root of negative number %g", 5.222)
	println()
	fmt.Println(errs)

}
