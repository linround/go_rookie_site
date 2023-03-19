package main

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"reflect"
)

type Values map[string][]string

func main() {
	f3()
}
func f2() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
}
func f3() {

	var vt = reflect.ValueOf(Values{})
	var it = reflect.ValueOf(45)
	fmt.Println(vt.Kind() == reflect.Map, "kind")
	fmt.Println(vt.Type())
	fmt.Println(reflect.TypeOf(Values{}))
	fmt.Println(vt)
	fmt.Println(it)
	fmt.Println(reflect.ValueOf(url.Values{}))
}

func f1() {

	var vt = reflect.TypeOf(Values{})
	var it = reflect.TypeOf(45)
	fmt.Println(it)
	fmt.Println(vt.String())
	fmt.Println(reflect.TypeOf(url.Values{}))
}
