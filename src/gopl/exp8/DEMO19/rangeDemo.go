package main

import "fmt"

func main() {
	var tmp = map[string]string{"name": "111", "age": "222"}
	for k, v := range tmp {
		fmt.Println(k, ":", v)
	}
	var arr = []string{"1", "2"}
	for id, v := range arr {
		fmt.Println(id, ":", v)
	}
}
