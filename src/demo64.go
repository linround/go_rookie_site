package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}
func (s stockPosition) getText() string {
	return s.ticker
}

type care struct {
	make  string
	model string
	price float32
}

func (c care) getValue() float32 {
	return c.price
}
func (c care) getText() string {
	return c.model
}

type valuable interface {
	text
	getValue() float32
}

type text interface {
	getText() string
}

func showValue(a valuable) {
	fmt.Println(a.getValue())
}

func showText(t valuable) {
	fmt.Println(t.getText())
}

// 从以下的应用程序中可以看出
// 在普通的程序中 结构体用来实现数据
// 接口通常用来提前定义相关的变量
func main() {
	o := new(stockPosition)
	o.count = 20
	o.sharePrice = 60
	o.ticker = "ticker"
	var oInt valuable
	oInt = o
	//	stockPosition{"gg", 100, 200}
	showValue(o)
	//o = care{"car", "m3", 200}
	showValue(o)
	showText(o)
	// 此处用来判断 接口对应的结构体类型
	u, ok := oInt.(*stockPosition)
	fmt.Println("u:", u, ok)
	if ok {
		fmt.Println("u:", u)
	}

}
