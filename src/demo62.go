package main

import "fmt"

type Log struct {
	msg string
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}
func (l *Log) Str() string {
	return l.msg
}

type Customer struct {
	Name string
	log  *Log
}

func (c *Customer) Log() *Log {
	return c.log
}

type Base struct {
}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (se Base) MoreMagic() {
	se.Magic()
	se.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo Magic")
}

func main() {
	//c := new(Customer)
	//c.Name = "BABA"
	//c.log = new(Log)
	//c.log.msg = "1111111"

	c := &Customer{"A", &Log{"222222"}}
	c.Log().Add("33333")
	fmt.Println(c.Log())

	v := new(Voodoo)
	// 以下打印顺序说明，结构体会调用离自己最近的方法
	v.Magic()     // voodoo magic
	v.MoreMagic() // base magic
}
