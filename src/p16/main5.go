package main

type struct1 struct {
	name string
}

// 这是一个指针
var ms = new(struct1)

// 通常建议使用构建函数来初始化结构体
// 不推荐使用声明的方式进行构建结构体
func newStruct1(name string) *struct1 {
	return &struct1{
		name: name,
	}
}

func main() {
	ms = &struct1{name: "linyuan"}
}
