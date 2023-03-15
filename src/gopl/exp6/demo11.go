package main

import "fmt"

func main() {
	res := fmt.Errorf("%s \n %s", "dddddddddddddddddd", "cccccccccccc")
	fmt.Println(res)
	// Errno(2) 把2进行类型转换，转换为Errno类型
	// Errno 实现了对应的Error 方法

	// error 是一个 接口类型 ,这个接口类型具有Error 方法
	// 所以可以将一个已经实现了Error 方法的类型Errno(2) 赋值给一个接口类型 err
	var err error = Errno{"a": "aaaaaaa", "b": "bbbbb", "c": "cccccccccc"}
	fmt.Println(err.Error())
}

type Errno map[string]string

func (e Errno) Error() string {
	return "myError"
}
