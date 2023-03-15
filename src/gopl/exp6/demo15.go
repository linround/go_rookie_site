package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	test()
	t2()
}

func test() {
	fmt.Println(strings.Contains("你好", "你"))
	fmt.Println(strings.Contains("你好", "ni"))
}
func t2() {
	_, err := os.Open("/no/such/file")
	fmt.Println(os.IsNotExist(err))
	fmt.Println(err, "8888888888")

	fmt.Printf("%#v\n", err)

}
