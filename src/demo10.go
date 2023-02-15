package main

import "fmt"

func print9x() {
	for m := 1; m < 10; m++ {
		for n := 1; n <= m; n++ {
			fmt.Printf("%dx%d=%d ", n, m, m*n)
		}
		fmt.Println("")
	}
}

func gotoTag() {
	for i := 1; i <= 9; i++ {
		n := 1
	Loop:
		if n <= i {
			fmt.Printf("%dx%d=%d ", i, n, i*n)
			n++
			goto Loop
		} else {
			n = 1
			fmt.Println("")
		}

	}
}

func main() {
	print9x()
	gotoTag()
}
