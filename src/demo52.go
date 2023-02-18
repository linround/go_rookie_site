package main

import (
	"fmt"
	"regexp"
)

func main() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"
	ok, v := regexp.Match(pat, []byte(searchIn))
	fmt.Println(ok, v)
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)

}
