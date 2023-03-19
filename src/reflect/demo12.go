package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3})) // "true"
	fmt.Println(Equal([]string{"foo"}, []string{"bar"})) // "false"
	//fmt.Println(Equal([]string(nil), []string{}))             // "true"
	//fmt.Println(Equal(map[string]int(nil), map[string]int{})) // "true"

}

// Equal reports whether x and y are deeply equal.
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	// ...cycle check omitted (shown later)...

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()
	case reflect.String:
		return x.String() == y.String()

	// ...numeric cases omitted for brevity...

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()
	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)
	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			flag := equal(x.Index(i), y.Index(i), seen)
			if !flag {
				return false
			}
		}
		return true

		// ...struct and map cases omitted for brevity...
	}
	panic("unreachable")
}
