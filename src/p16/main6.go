package main

import "fmt"

// 关于如何监测一个值是否实现了某个接口

type stringer interface {
	value() string
}

func testStringer(v interface{}) {
	if v, ok := v.(stringer); ok {
		fmt.Println("这个值实现了该接口", v)
	}
}

type str struct {
}

func (s *str) value() string {
	return "实现了value方法"
}

// 使用接口实现一个类型分类

func classF(items []interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Println("布尔值", i, x)
		case int:
			fmt.Println("数字", i, x)
		case stringer:
			fmt.Println("其他类型", i, x)
		case string:
			fmt.Println("字符串", i, x)
		}
	}
}

func main() {
	var s = new(str)
	testStringer(s)
	types := []interface{}{1, 2, s, true, false, "1554"}
	classF(types)

}
