package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert2 Employee
var dilbert3 *Employee

func main() {
	var dilbert Employee
	//这里的赋值是一个 取值
	dilbert2 = dilbert
	// 这里的赋值是一个 取址
	dilbert3 = &dilbert
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code
	fmt.Println(dilbert2.Salary)
	fmt.Println(dilbert3.Salary)
}
