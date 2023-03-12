package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println(string(65))      // "A", not "65"
	fmt.Println(string(0x4eac))  // "京"
	fmt.Println(string(0xFFFD))  // ?
	fmt.Println(string(1234567)) // "?"

	fmt.Printf("%c", '\uFFFD')
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c    \n", i, r)
		i += size
	}

}
