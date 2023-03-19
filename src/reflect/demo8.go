package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type iii int

func (t iii) dd() int {
	return 0
}

// Print prints the method set of the value x.
func Print(x interface{}) {
	v := reflect.ValueOf(x)
	t := v.Type()
	fmt.Printf("type %s\n", t)

	for i := 0; i < v.NumMethod(); i++ {
		methType := v.Method(i).Type()
		fmt.Printf("func (%s) ---  %s -- %s --%s\n",
			t,
			t.Method(i).Name,
			methType.String(),
			strings.TrimPrefix(methType.String(), "func"))
	}
}

func main() {
	var i iii
	i = 45
	Print(i)
	time.Hour.Hours()
}
