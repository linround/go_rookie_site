package main

import "fmt"

// 由于接口可能实现在很多种类型上
// 所以需要检测对应接口变量被所属的数据类型

// 同意有时候也需要测试某个值是否实现了某个接口
//

// 定义一个list结构体
type list []int

func (l list) len() int {
	return len(l)
}
func (l *list) append(v int) {
	*l = append(*l, v)
}

// 定义一个接口
type appender interface {
	append(int2 int)
}

// countInto 是需要一个appender
// 而 a的方法append 只定义在指针上
func countInto(a appender, start, end int) {
	for i := start; i < end; i++ {
		a.append(i)
	}
}

type lener interface {
	len() int
}

// 定义一个函数
func longEnough(l lener) bool {
	return l.len()*10 > 42
}

func main() {
	var lst list
	// list的append方法是定义在指针上面的

	countInto(&lst, 1, 10)

	if longEnough(lst) {
		fmt.Println("111111111111")
	}

	// 指针会被自动解引用
	// 值不会自动解引用 所以需要添加&号
	//pLst := list{1, 2}
	//countInto(&pLst, 1, 10)
	pLst := new(list)
	countInto(pLst, 1, 10)
	if longEnough(pLst) {
		fmt.Println("2222222222")
	}
}
