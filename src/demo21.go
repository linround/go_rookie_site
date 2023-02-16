package main

import "fmt"

// 定义一个divideError 结构体

type DivideError struct {
	dividee int
	divider int
}

func (de DivideError) Error() string {
	str := "这是定义的错误返回字符串: %d"
	return fmt.Sprintf(str, de.divider)
}

func Divide(varDividee int, varDivider int) (result int, err string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		err = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	result, err := Divide(100, 10)
	if err == "" {
		fmt.Println("正确", result)
	}
	_, error := Divide(100, 0)
	if error != "" {
		// 出现了错误
		fmt.Println("可以不整除")
		fmt.Println(error)
	} else {
		fmt.Println(error)
	}
}
