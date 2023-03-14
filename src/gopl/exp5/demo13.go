package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}}
	m.Add("item", "1")
	m.Add("item", "2")
	m.Add("item", "3")
	for _, v := range m["item"] {
		fmt.Println(v)
	}
	fmt.Println(m["item"])
	fmt.Println(m.Get("item"))
	m = nil
	fmt.Println(m.Get("item")) // ""
	fmt.Println("**1")
}
