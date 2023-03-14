package main

import (
	"fmt"
	"time"
)

var t time.Time

func main() {
	fmt.Println(t.Location())
	fmt.Println(t.Date())
}
