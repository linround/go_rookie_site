package main

import (
	"fmt"
	"math"
	"strconv"
)

type Expr interface {
	// Eval 接口定义的方法
	Eval(env Env) float64
	String() string
}

// Var 变量
type Var string

// literal 常量
type literal float64

// 实现 常量和变量的Eval方法

func (v Var) Eval(env Env) float64 {
	return env[v]
}
func (l literal) Eval(Env) float64 {
	return float64(l)
}
func (v Var) String() string {
	return "变量:" + string(v)
}
func (l literal) String() string {
	return "常量:" + strconv.FormatFloat(float64(l), 'f', -1, 64)

}

// 一元符号
type unary struct {
	op rune
	x  Expr
}

// 二元符号
type binary struct {
	op   rune
	x, y Expr
}

// 实现对运算符的操作

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("没有支持的运算符，%q", u.op))
}
func (u unary) String() string {
	return "(操作符号:" + strconv.QuoteRuneToASCII(u.op) + " | " +
		u.x.String() + ")"

}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("没有支持的运算符，%q", b.op))
}
func (b binary) String() string {
	return "(" + b.x.String() + " | 操作符号:" +
		strconv.QuoteRuneToASCII(b.op) + " | " + b.y.String() + ")"

}

// 函数调用表达式
type call struct {
	fn   string
	args []Expr
}

// 实现对函数的调用

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("没有支持的函数，%q", c.fn))

}
func (c call) String() string {
	var args string
	for i, v := range c.args {
		args += v.String()
		if i < len(c.args)-1 {
			args += ", "
		}
	}
	//fmt.Println(args)
	return "函数:" + c.fn + "(" + args + ")"

}

// Env 映射对应的变量值
type Env map[Var]float64

func main() {
	env := Env{"x": 3, "y": 4}
	xy := unary{
		op: '-',
		x:  Var("y"),
	}

	add := binary{
		op: '+',
		x:  Var("x"),
		y:  Var("y"),
	}

	// 乘
	mul := binary{
		op: '*',
		x:  Var("x"),
		y:  Var("y"),
	}

	// pow
	pow := call{
		fn:   "pow",
		args: []Expr{Var("x"), Var("y")},
	}

	fmt.Println("xy:", add.Eval(env))
	fmt.Println(xy.Eval(env))
	fmt.Println(mul.Eval(env))
	fmt.Println(pow.Eval(env))

	var val Expr = Var("x")
	fmt.Println(val)
	val = literal(234.323233)
	fmt.Println(val)
	fmt.Println(xy)
	fmt.Println(add)
	fmt.Println(mul)
	fmt.Println(pow)

	wawa := Env{"q": 1, "w": 2, "e": 3, "x": 3, "y": 4}
	fmt.Println("变量列表：", wawa)
	taowa := call{
		fn: "pow",
		args: []Expr{
			add,
			call{
				fn: "sqrt",
				args: []Expr{
					binary{
						op: '+',
						x:  Var("e"),
						y:  Var("q"),
					},
				},
			},
		},
	}
	fmt.Println(taowa)            // 函数:pow((变量:x | 操作符号:'+' | 变量:y), 函数:sqrt((变量:e | 操作符号:'+' | 变量:q)))
	fmt.Println(taowa.Eval(wawa)) // 49
}
