package main

import (
	"fmt"
	"net/url"
	"strconv"
)

type t struct {
}
type stringer interface {
	String() string
}

func main() {
	s := Sprint(Values{})
	fmt.Println(s)
}

// 虽然这个Values与url.Values一模一样。但是也没办法识别出是一样的结构
// 没有办法检测未知类型的表示方式。所以我们需要反射来处理这个问题

type Values map[string][]string

func Sprint(x interface{}) string {

	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	// ...similar cases for int16, uint32, and so on...
	case bool:
		if x {
			return "true"
		}
		return "false"
	case url.Values:
		return "url.Values"
	default:
		// array, chan, func, map, pointer, slice, struct
		return "???"
	}
}
